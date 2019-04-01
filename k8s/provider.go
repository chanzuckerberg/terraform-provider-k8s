package k8s

import (
	"bytes"

	"github.com/hashicorp/terraform/helper/schema"
	restclient "k8s.io/client-go/rest"
)

func Provider() *schema.Provider {
	s := &schema.Provider{}

	providerConfigure := func(d *schema.ResourceData) (interface{}, error) {
		cfg := &restclient.Config{}
		if v, ok := d.GetOk("host"); ok {
			cfg.Host = v.(string)
		}
		if v, ok := d.GetOk("username"); ok {
			cfg.Username = v.(string)
		}
		if v, ok := d.GetOk("password"); ok {
			cfg.Password = v.(string)
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
		SetK8SSingletonClientConfig(cfg)
		k8sConfig := K8SConfig_Singleton()
		k8sConfig.Namespace = d.Get("namespace").(string)

		// HACK set these here since we depend on the
		// provier being configured already
		s.ResourcesMap = BuildResourcesMap()
		s.DataSourcesMap = BuildDataSourcesMap()

		return k8sConfig, nil
	}

	s.Schema = map[string]*schema.Schema{
		"namespace": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "default",
			Description: "Default namespace",
		},
		"host": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_HOST", ""),
			Description: "The hostname (in form of URI) of Kubernetes master.",
		},
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_USER", ""),
			Description: "The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_PASSWORD", ""),
			Description: "The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
		},
		"client_certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CLIENT_CERT_DATA", ""),
			Description: "PEM-encoded client certificate for TLS authentication.",
		},
		"client_key": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CLIENT_KEY_DATA", ""),
			Description: "PEM-encoded client certificate key for TLS authentication.",
		},
		"cluster_ca_certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CLUSTER_CA_CERT_DATA", ""),
			Description: "PEM-encoded root certificates bundle for TLS authentication.",
		},
		"config_path": {
			Type:     schema.TypeString,
			Optional: true,
			DefaultFunc: schema.MultiEnvDefaultFunc(
				[]string{
					"KUBE_CONFIG",
					"KUBECONFIG",
				},
				"~/.kube/config"),
			Description: "Path to the kube config file, defaults to ~/.kube/config",
		},
		"config_context": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CTX", ""),
		},
		"config_context_auth_info": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CTX_AUTH_INFO", ""),
			Description: "",
		},
		"config_context_cluster": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CTX_CLUSTER", ""),
			Description: "",
		},
		"token": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_TOKEN", ""),
			Description: "Token to authenticate an service account",
		},
	}

	s.ConfigureFunc = providerConfigure
	return s
}
