package provider

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func newDynamicClient(model KubernetesModel) (dynamic.Interface, error) {
	config, err := createKubernetesConfig(model)
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(config)
}

func newDiscoveryClient(model KubernetesModel) (discovery.DiscoveryInterface, error) {
	config, err := createKubernetesConfig(model)
	if err != nil {
		return nil, err
	}
	return discovery.NewDiscoveryClientForConfig(config)
}

// Reference: https://github.com/hashicorp/terraform-provider-kubernetes/blob/main/kubernetes/provider.go#L494
func createKubernetesConfig(model KubernetesModel) (*rest.Config, error) {
	overrides := &clientcmd.ConfigOverrides{}
	loader := &clientcmd.ClientConfigLoadingRules{}

	// Handle kubeconfig and context
	var configPaths []string
	if model.ConfigPath.ValueString() != "" {
		configPaths = []string{model.ConfigPath.ValueString()}
	} else if len(model.ConfigPaths) > 0 {
		for _, path := range model.ConfigPaths {
			configPaths = append(configPaths, path.ValueString())
		}
	} else if v := os.Getenv("KUBE_CONFIG_PATHS"); v != "" {
		configPaths = filepath.SplitList(v)
	}

	if len(configPaths) > 0 {
		var expandedPaths []string
		for _, p := range configPaths {
			path, err := homedir.Expand(p)
			if err != nil {
				return nil, err
			}

			expandedPaths = append(expandedPaths, path)
		}

		if len(expandedPaths) == 1 {
			loader.ExplicitPath = expandedPaths[0]
		} else {
			loader.Precedence = expandedPaths
		}

		// Handle context in kubeconfig
		ctxSuffix := "; default context"
		if model.ConfigContext.ValueString() != "" {
			ctxSuffix = "; overriden context"
			overrides.CurrentContext = model.ConfigContext.ValueString()
			ctxSuffix += fmt.Sprintf("; config ctx: %s", overrides.CurrentContext)
			overrides.Context = api.Context{}
		}
	}

	if model.Insecure.ValueBool() {
		overrides.ClusterInfo.InsecureSkipTLSVerify = model.Insecure.ValueBool()
	}
	if model.ClusterCACertificate.ValueString() != "" {
		overrides.ClusterInfo.CertificateAuthorityData = bytes.NewBufferString(model.ClusterCACertificate.ValueString()).Bytes()
	}
	if model.ClientCertificate.ValueString() != "" {
		overrides.AuthInfo.ClientCertificateData = bytes.NewBufferString(model.ClientCertificate.ValueString()).Bytes()
	}
	if model.Host.ValueString() != "" {
		hasCA := len(overrides.ClusterInfo.CertificateAuthorityData) != 0
		hasCert := len(overrides.AuthInfo.ClientCertificateData) != 0
		defaultTLS := hasCA || hasCert || overrides.ClusterInfo.InsecureSkipTLSVerify
		host, _, err := rest.DefaultServerURL(model.Host.ValueString(), "", schema.GroupVersion{}, defaultTLS)
		if err != nil {
			return nil, fmt.Errorf("failed to parse host: %w", err)
		}

		overrides.ClusterInfo.Server = host.String()
	}
	if username := model.Username.ValueString(); username != "" {
		overrides.AuthInfo.Username = username
	}
	if password := model.Password.ValueString(); password != "" {
		overrides.AuthInfo.Password = password
	}
	if clientKey := model.ClientKey.ValueString(); clientKey != "" {
		overrides.AuthInfo.ClientKeyData = bytes.NewBufferString(clientKey).Bytes()
	}
	if token := model.Token.ValueString(); token != "" {
		overrides.AuthInfo.Token = token
	}

	// Create configuration
	cc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)
	return cc.ClientConfig()
}
