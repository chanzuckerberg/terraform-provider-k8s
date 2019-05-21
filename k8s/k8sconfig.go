package k8s

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	tfSchema "github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	tfHomedir "github.com/mitchellh/go-homedir"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	diskcached "k8s.io/client-go/discovery/cached/disk"
	"k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
	"k8s.io/kube-openapi/pkg/util/proto"
)

// NewK8SConfig gets a K8SConfig object configured from the given settings
// Functionality extracted from a combination of terraform-provider-kubernetes/kubernetes/provider.go
// and k8s.io/cli-runtime/pkg/genericclioptions/config_flags.go to be able to pass in our own
// restclient.Config
func NewK8SConfig(d *tfSchema.ResourceData) (*K8SConfig, error) {
	var cfg *restclient.Config
	var err error
	if d.Get("load_config_file").(bool) {
		// Config file loading
		cfg, err = tryLoadingConfigFile(d)
	}

	if err != nil {
		return nil, err
	}
	if cfg == nil {
		cfg = &restclient.Config{}
	}

	// Overriding with static configuration
	cfg.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraform.VersionString())

	if v, ok := d.GetOk("host"); ok {
		cfg.Host = v.(string)
	}
	if v, ok := d.GetOk("username"); ok {
		cfg.Username = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		cfg.Password = v.(string)
	}
	if v, ok := d.GetOk("insecure"); ok {
		cfg.Insecure = v.(bool)
	}
	if v, ok := d.GetOk("cluster_ca_certificate"); ok {
		cfg.CAData = bytes.NewBufferString(v.(string)).Bytes()
	}
	if v, ok := d.GetOk("client_certificate"); ok {
		cfg.CertData = bytes.NewBufferString(v.(string)).Bytes()
	}
	if v, ok := d.GetOk("client_key"); ok {
		cfg.KeyData = bytes.NewBufferString(v.(string)).Bytes()
	}
	if v, ok := d.GetOk("token"); ok {
		cfg.BearerToken = v.(string)
	}

	if v, ok := d.GetOk("exec"); ok {
		exec := &clientcmdapi.ExecConfig{}
		if spec, ok := v.([]interface{})[0].(map[string]interface{}); ok {
			exec.APIVersion = spec["api_version"].(string)
			exec.Command = spec["command"].(string)
			exec.Args = expandStringSlice(spec["args"].([]interface{}))
			for kk, vv := range spec["env"].(map[string]interface{}) {
				exec.Env = append(exec.Env, clientcmdapi.ExecEnvVar{Name: kk, Value: vv.(string)})
			}
		} else {
			return nil, fmt.Errorf("Failed to parse exec")
		}
		cfg.ExecProvider = exec
	}

	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// TODO(mbarrien): Switch to tfHomedir.Expand("~/.kube/http-cache") ?
	var httpCacheDir = filepath.Join(homedir.HomeDir(), ".kube", "http-cache")

	// The more groups you have, the more discovery requests you need to make.
	// given 25 groups (our groups + a few custom resources) with one-ish version each, discovery needs to make 50 requests
	// double it just so we don't end up here again for a while.  This config is only used for discovery.
	cfg.Burst = 100
	// TODO(mbarrien): Switch to tfHomedir.Expand("~/.kube/cache/discovery") ?
	discoveryCacheDir := computeDiscoverCacheDir(filepath.Join(homedir.HomeDir(), ".kube", "cache", "discovery"), cfg.Host)
	discoveryClient, err := diskcached.NewCachedDiscoveryClientForConfig(cfg, discoveryCacheDir, httpCacheDir, time.Duration(10*time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)
	expander := restmapper.NewShortcutExpander(mapper, discoveryClient)

	modelsMap := buildModelsMap(discoveryClient)
	tfSchemasMap := buildTFSchemasMap(modelsMap)

	return &K8SConfig{
		RESTMapper:      expander,
		DynamicClient:   dynamicClient,
		DiscoveryClient: discoveryClient,
		ModelsMap:       modelsMap,
		TFSchemasMap:    tfSchemasMap,
	}, nil
}

// overlyCautiousIllegalFileCharacters matches characters that *might* not be supported.  Windows is really restrictive, so this is really restrictive
var overlyCautiousIllegalFileCharacters = regexp.MustCompile(`[^(\w/\.)]`)

// computeDiscoverCacheDir takes the parentDir and the host and comes up with a "usually non-colliding" name.
func computeDiscoverCacheDir(parentDir, host string) string {
	// strip the optional scheme from host if its there:
	schemelessHost := strings.Replace(strings.Replace(host, "https://", "", 1), "http://", "", 1)
	// now do a simple collapse of non-AZ09 characters.  Collisions are possible but unlikely.  Even if we do collide the problem is short lived
	safeHost := overlyCautiousIllegalFileCharacters.ReplaceAllString(schemelessHost, "_")
	return filepath.Join(parentDir, safeHost)
}

func tryLoadingConfigFile(d *tfSchema.ResourceData) (*restclient.Config, error) {
	path, err := tfHomedir.Expand(d.Get("config_path").(string))
	if err != nil {
		return nil, err
	}

	loader := &clientcmd.ClientConfigLoadingRules{
		ExplicitPath: path,
	}

	overrides := &clientcmd.ConfigOverrides{}
	ctxSuffix := "; default context"

	ctx, ctxOk := d.GetOk("config_context")
	authInfo, authInfoOk := d.GetOk("config_context_auth_info")
	cluster, clusterOk := d.GetOk("config_context_cluster")
	if ctxOk || authInfoOk || clusterOk {
		ctxSuffix = "; overriden context"
		if ctxOk {
			overrides.CurrentContext = ctx.(string)
			ctxSuffix += fmt.Sprintf("; config ctx: %s", overrides.CurrentContext)
			log.Printf("[DEBUG] Using custom current context: %q", overrides.CurrentContext)
		}

		overrides.Context = clientcmdapi.Context{}
		if authInfoOk {
			overrides.Context.AuthInfo = authInfo.(string)
			ctxSuffix += fmt.Sprintf("; auth_info: %s", overrides.Context.AuthInfo)
		}
		if clusterOk {
			overrides.Context.Cluster = cluster.(string)
			ctxSuffix += fmt.Sprintf("; cluster: %s", overrides.Context.Cluster)
		}
		log.Printf("[DEBUG] Using overidden context: %#v", overrides.Context)
	}

	cc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)
	cfg, err := cc.ClientConfig()
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok && os.IsNotExist(pathErr.Err) {
			log.Printf("[INFO] Unable to load config file as it doesn't exist at %q", path)
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to load config (%s%s): %s", path, ctxSuffix, err)
	}

	log.Printf("[INFO] Successfully loaded config file (%s%s)", path, ctxSuffix)
	return cfg, nil
}

type K8SConfig struct {
	RESTMapper      meta.RESTMapper
	DynamicClient   dynamic.Interface
	DiscoveryClient discovery.CachedDiscoveryInterface
	cache           sync.Map
	countdownLatch  sync.Map
	mutex           sync.Mutex
	ModelsMap       map[schema.GroupVersionKind]proto.Schema
	TFSchemasMap    map[string]*tfSchema.Schema
}

func (this *K8SConfig) Get(name string, getOption metav1.GetOptions, gvk *schema.GroupVersionKind, namespace string) (*unstructured.Unstructured, error) {
	id := CreateId(namespace, gvk.Kind, name)
	//log.Println("Get name:", name, "gvk:", gvk, "id:", id)
	//try getting without lock first
	if item, ok := this.cache.Load(id); ok {
		return item.(*unstructured.Unstructured), nil
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	//try again with lock
	if item, ok := this.cache.Load(id); ok {
		return item.(*unstructured.Unstructured), nil
	}

	//get one until countdown countdownLatch
	if latch, ok := this.countdownLatch.Load(gvk); ok {
		if latch.(int) == 0 {
			//get all
			if res, err := this.GetAll(gvk, namespace); err != nil {
				return nil, err
			} else {
				for _, item := range res.Items {
					itemName, _, _ := unstructured.NestedString(item.Object, "metadata", "name")
					itemId := CreateId(namespace, gvk.Kind, itemName)
					this.cache.Store(itemId, item.DeepCopy())
				}
			}
		} else {
			//count down
			this.countdownLatch.Store(gvk, latch.(int)-1)
			one, err := this.GetOne(name, gvk, namespace)
			if err != nil {
				return nil, err
			}
			this.cache.Store(id, one)
		}
	} else {
		//todo: make countdown value configurable
		//start countdown
		this.countdownLatch.Store(gvk, 0)
		one, err := this.GetOne(name, gvk, namespace)
		if err != nil {
			return nil, err
		}
		this.cache.Store(id, one)
	}

	//try for the last time now that cache should be loaded
	if item, ok := this.cache.Load(id); ok {
		return item.(*unstructured.Unstructured), nil
	} else {
		return nil, errors.NewNotFound(schema.GroupResource{Group: gvk.Group, Resource: gvk.Kind}, name)
	}
}

func (this *K8SConfig) GetOne(name string, gvk *schema.GroupVersionKind, namespace string) (*unstructured.Unstructured, error) {
	RESTMapping, _ := this.RESTMapper.RESTMapping(schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}, gvk.Version)
	var resourceClient dynamic.ResourceInterface
	resourceClient = this.DynamicClient.Resource(RESTMapping.Resource)
	if namespace != "" {
		resourceClient = resourceClient.(dynamic.NamespaceableResourceInterface).Namespace(namespace)
	}
	res, err := resourceClient.Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	//log.Println("GetOne name:", name, "gvk:", gvk)
	return res, nil
}

func (this *K8SConfig) GetAll(gvk *schema.GroupVersionKind, namespace string) (*unstructured.UnstructuredList, error) {
	RESTMapping, _ := this.RESTMapper.RESTMapping(schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}, gvk.Version)
	var resourceClient dynamic.ResourceInterface
	resourceClient = this.DynamicClient.Resource(RESTMapping.Resource)
	if namespace != "" {
		resourceClient = resourceClient.(dynamic.NamespaceableResourceInterface).Namespace(namespace)
	}
	//log.Println("GetAll gvk:", gvk)
	return resourceClient.List(metav1.ListOptions{})
}

func (this *K8SConfig) ForEachAPIResource(callback func(metav1.APIResource, schema.GroupVersionKind)) {
	lists, _ := this.DiscoveryClient.ServerResources()
	for _, list := range lists {
		//log.Println("name:", group.Name, "group:", group.PreferredVersion.GroupVersion, "version:", group.PreferredVersion.Version)
		if len(list.APIResources) == 0 {
			continue
		}
		gv, _ := schema.ParseGroupVersion(list.GroupVersion)
		for _, apiResource := range list.APIResources {
			gvk, _ := this.RESTMapper.KindFor(schema.GroupVersionResource{
				Group:    gv.Group,
				Version:  gv.Version,
				Resource: apiResource.Kind,
			})
			callback(apiResource, gvk)
		}
	}
}

func buildModelsMap(DiscoveryClient discovery.DiscoveryInterface) map[schema.GroupVersionKind]proto.Schema {
	doc, err := DiscoveryClient.OpenAPISchema()
	if err != nil {
		log.Fatal(err)
	}
	models, _ := proto.NewOpenAPIData(doc)
	//get metadata from this for crd resources without spec
	nsModel := models.LookupModel("io.k8s.api.core.v1.Namespace").(*proto.Kind)
	modelsMap := map[schema.GroupVersionKind]proto.Schema{}
	for _, modelName := range models.ListModels() {
		model := models.LookupModel(modelName)
		if model == nil {
			//log.Println("No Model For ModelName:", modelName)
			continue
		}
		gvkList := parseGroupVersionKind(model)
		for _, gvk := range gvkList {
			if len(gvk.Kind) > 0 && !IsSkipKind(gvk.Kind) {
				//check for crd resources without spec
				if _, isMap := model.(*proto.Map); isMap {
					modelsMap[gvk] = &proto.Kind{
						BaseSchema: model.(*proto.Map).BaseSchema,
						Fields: map[string]proto.Schema{
							"metadata": nsModel.Fields["metadata"],
							"spec":     model.(*proto.Map).SubType, //must be type Arbitrary
						},
					}
				} else {
					modelsMap[gvk] = model
				}
			}
		}
	}
	return modelsMap
}

func buildTFSchemasMap(modelsMap map[schema.GroupVersionKind]proto.Schema) map[string]*tfSchema.Schema {
	schemasMap := map[string]*tfSchema.Schema{}
	for gvk, model := range modelsMap {
		resourceKey := ResourceKey(gvk.Group, gvk.Version, gvk.Kind)
		//log.Println("gvk:", gvk, "resource:", resourceKey)
		if _, hasKey := schemasMap[resourceKey]; hasKey {
			//dups
			//sometimes resources appear more than once
			continue
		}

		schemaVisitor := NewK8S2TFSchemaVisitor(resourceKey)
		model.Accept(schemaVisitor)
		schemasMap[resourceKey] = &schemaVisitor.Schema
	}
	return schemasMap
}

const groupVersionKindExtensionKey = "x-kubernetes-group-version-kind"

func parseGroupVersionKind(s proto.Schema) []schema.GroupVersionKind {
	extensions := s.GetExtensions()

	gvkListResult := []schema.GroupVersionKind{}

	// Get the extensions
	gvkExtension, ok := extensions[groupVersionKindExtensionKey]
	if !ok {
		return []schema.GroupVersionKind{}
	}

	// gvk extension must be a list of at least 1 element.
	gvkList, ok := gvkExtension.([]interface{})
	if !ok {
		return []schema.GroupVersionKind{}
	}

	for _, gvk := range gvkList {
		// gvk extension list must be a map with group, version, and
		// kind fields
		gvkMap, ok := gvk.(map[interface{}]interface{})
		if !ok {
			continue
		}
		group, ok := gvkMap["group"].(string)
		if !ok {
			continue
		}
		version, ok := gvkMap["version"].(string)
		if !ok {
			continue
		}
		kind, ok := gvkMap["kind"].(string)
		if !ok {
			continue
		}

		gvkListResult = append(gvkListResult, schema.GroupVersionKind{
			Group:   group,
			Version: version,
			Kind:    kind,
		})
	}

	return gvkListResult
}

func expandStringSlice(s []interface{}) []string {
	result := make([]string, len(s), len(s))
	for k, v := range s {
		// Handle the Terraform parser bug which turns empty strings in lists to nil.
		if v == nil {
			result[k] = ""
		} else {
			result[k] = v.(string)
		}
	}
	return result
}
