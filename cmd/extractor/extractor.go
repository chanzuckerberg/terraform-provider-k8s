package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/kube-openapi/pkg/util/proto"

	"github.com/mingfang/terraform-provider-k8s/k8s"
)

func main() {
	var filename, namespace, kind, name, dir, url, kubeconfig, cacheDir string
	var isImport bool

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&filename, "filename", "", "name of file to extract")
	f.StringVar(&filename, "file", "", "name of file to extract")
	f.StringVar(&filename, "f", "", "name of file to extract")
	f.StringVar(&dir, "dir", "", "destination directory")
	f.StringVar(&url, "url", "", "source url")
	f.StringVar(&namespace, "namespace", "default", "namespace of resources to extract")
	f.StringVar(&namespace, "ns", "default", "namespace of resources to extract")
	f.StringVar(&kind, "kind", "", "kind of resources to extract")
	f.StringVar(&name, "name", "", "name of resources to extract")
	f.StringVar(&kubeconfig, "kubeconfig", "", "path to kubeconfig file")
	f.StringVar(&kubeconfig, "k", "", "path to kubeconfig file")
	f.StringVar(&cacheDir, "cachedir", "", "path to kubernetes cache directory")
	f.BoolVar(&isImport, "import", false, "automatically import resources")
	f.Parse(os.Args[1:])

	k8sConfig, err := k8s.NewK8SConfigFromKubeconfig(kubeconfig, cacheDir)
	if err != nil {
		log.Println(err)
		return
	}
	if url != "" {
		extractURL(k8sConfig, url, kind, dir)
	} else if filename != "" {
		extractFile(k8sConfig, filename, kind, dir)
	} else if namespace != "" || kind != "" || name != "" {
		extractCluster(k8sConfig, namespace, kind, name, isImport, dir)
	} else {
		fmt.Println("Usage: -filename <name of file to extract>")
		fmt.Println("Usage: -namespace <namespace> -kind <kind> -name <name>, blank means all")
	}

}
func extractURL(k8sConfig *k8s.K8SConfig, url string, kind string, dir string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	extractYamlBytes(k8sConfig, body, kind, dir)
}

func extractFile(k8sConfig *k8s.K8SConfig, filename string, kind string, dir string) {
	yamlBytes, yamlErr := ioutil.ReadFile(filename)
	if yamlErr != nil {
		log.Fatal(yamlErr)
	}
	extractYamlBytes(k8sConfig, yamlBytes, kind, dir)
}

func extractYamlBytes(k8sConfig *k8s.K8SConfig, yamlBytes []byte, kindFilter string, dir string) {
	modelsMap := k8sConfig.ModelsMap

	decoder := yaml.NewYAMLToJSONDecoder(bytes.NewReader(yamlBytes))
	var decodeErr error
	for {
		var object map[string]interface{}
		decodeErr = decoder.Decode(&object)
		if decodeErr != nil {
			break
		}
		if len(object) == 0 {
			continue
		}
		//log.Println(object)
		kind := object["kind"].(string)
		var items []interface{}
		if kind == "List" {
			items = object["items"].([]interface{})
		} else {
			items = []interface{}{object}
		}

		for _, each := range items {
			item := each.(map[string]interface{})
			kind = item["kind"].(string)
			if kindFilter != "" && strings.ToLower(kindFilter) != strings.ToLower(kind) {
				//log.Println("Skip kind:", kind)
				continue
			}

			group, version, _ := k8s.SplitGroupVersion(item["apiVersion"].(string))
			resourceKey := k8s.ResourceKey(group, version, kind)
			gvk, _ := k8sConfig.RESTMapper.KindFor(schema.GroupVersionResource{
				Group:    group,
				Version:  version,
				Resource: kind,
			})
			model := modelsMap[gvk]
			if model == nil {
				log.Println("No Model For:", kind, gvk)
				continue
			}
			saveK8SasTF(item, model, resourceKey, gvk, dir)
		}
	}
	if decodeErr != nil && decodeErr != io.EOF {
		log.Println(decodeErr)
	}
}

func extractCluster(k8sConfig *k8s.K8SConfig, namespace, kind, name string, isImport bool, dir string) {
	systemNamePattern := regexp.MustCompile(`^system:`)
	var dupDetector = map[string]struct{}{}
	resourceVerbs := []string{"create", "get"}

	k8sConfig.ForEachAPIResource(func(apiResource metav1.APIResource, gvk schema.GroupVersionKind) {
		if kind != "" && strings.ToLower(kind) != strings.ToLower(apiResource.Kind) {
			//log.Println("Skip kind:", gvk.Group, gvk.Version, gvk.Kind)
			return
		}
		if len(apiResource.Verbs) > 0 && !sets.NewString(apiResource.Verbs...).HasAll(resourceVerbs...) {
			return
		}
		model := k8sConfig.ModelsMap[gvk]
		if model == nil {
			//log.Println("No Model For:", gvk)
			return
		}
		resourceKey := k8s.ResourceKey(gvk.Group, gvk.Version, gvk.Kind)
		//log.Println("gvk:", gvk, "resource:", resourceKey)

		//todo: handle dup deploys that comes in as both beta1 and v1
		//todo: prevent dup, some are repeated with category = "all"
		if _, hasKey := dupDetector[resourceKey]; hasKey {
			return
		}
		dupDetector[resourceKey] = struct{}{}

		//todo: add exclude option
		if apiResource.Kind == "Pod" || apiResource.Kind == "ReplicaSet" || apiResource.Kind == "Endpoints" || apiResource.Kind == "APIService" {
			//log.Println("Skip:", apiResource.Kind)
			return
		}

		RESTMapping, _ := k8sConfig.RESTMapper.RESTMapping(schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}, gvk.Version)
		//todo: get namespace from command line
		var resourceClient dynamic.ResourceInterface
		resourceClient = k8sConfig.DynamicClient.Resource(RESTMapping.Resource)
		if apiResource.Namespaced {
			resourceClient = resourceClient.(dynamic.NamespaceableResourceInterface).Namespace(namespace)
		}
		res, err := resourceClient.List(metav1.ListOptions{})
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println("read res:", res)

		//skip existing resources
		if isImport {
			execCommand("terraform", []string{"init"})
		}
		stateList, execErr := execCommand("terraform", []string{"state", "list"})
		if execErr != nil {
			log.Println(string(stateList))
		}

		for _, item := range res.Items {
			itemName, _, _ := unstructured.NestedString(item.Object, "metadata", "name")
			if (name != "" && name != itemName) || systemNamePattern.MatchString(itemName) {
				continue
			}
			saveK8SasTF(item.Object, model, resourceKey, gvk, dir)

			if !isImport {
				continue
			}

			//import
			stateName := resourceKey + "." + k8s.ToSnake(itemName)
			if strings.Contains(stateList, stateName) {
				log.Println("skip import:", stateName)
				continue
			}
			var id string
			if apiResource.Namespaced {
				id = k8s.CreateId(namespace, gvk.Kind, itemName)
			} else {
				id = k8s.CreateId("", gvk.Kind, itemName)
			}
			log.Printf("terraform import %s.%s %s\n", resourceKey, itemName, id)
			if cmdOut, execErr := execCommand("terraform", []string{"import", stateName, id}); execErr != nil {
				log.Println(string(cmdOut))
			}
		}
	})
}

func execCommand(cmd string, args []string) (string, error) {
	command := exec.Command(cmd, args...)
	command.Env = append(os.Environ(), "TF_LOG=")
	cmdOut, execErr := command.CombinedOutput()
	return string(cmdOut), execErr
}

func saveK8SasTF(itemObject map[string]interface{}, model proto.Schema, resourceKey string, gvk schema.GroupVersionKind, dir string) {
	var buf bytes.Buffer
	name, _, _ := unstructured.NestedString(itemObject, "metadata", "name")
	name = k8s.ToSnake(name)

	visitor := NewK8S2TFPrintVisitor(&buf, fmt.Sprintf("resource \"%s\" \"%s\"", resourceKey, name), resourceKey, itemObject, 0, false)
	model.Accept(visitor)

	filename := name + "-" + k8s.ToSnake(gvk.Kind) + ".tf"
	if dir != "" {
		filename = dir + "/" + filename
	}
	log.Println(filename)
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		log.Fatal("WriteFile err:", err)
	}
	//fmt
	if cmdOut, execErr := execCommand("terraform", []string{"fmt", filename}); execErr != nil {
		log.Println(string(cmdOut))
	}

}
