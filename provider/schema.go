package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var kustomizeProviderSchema = schema.Schema{
	Description: `This Kustomization provider is used to build Kubernetes manifests using Kustomization.
Although there are existing providers, this provider gives you the best DX and mitigate the datasource's read when apply issue.`,
	Attributes: map[string]schema.Attribute{
		"kubernetes": schema.SingleNestedAttribute{
			Description: "Kubernetes configuration used in `kustomize_apply",
			Optional:    true,
			Attributes:  kubernetesAttributes,
		},
	},
}

var kubernetesAttributes = map[string]schema.Attribute{
	"config_path": schema.StringAttribute{
		Description: "A path to a kube config file.",
		Optional:    true,
	},
	"config_paths": schema.ListAttribute{
		ElementType: types.StringType,
		Description: "A list of paths to the kube config files.",
		Optional:    true,
	},
	"host": schema.StringAttribute{
		Description: "The hostname (in form of URI) of the Kubernetes API.",
		Optional:    true,
	},
	"username": schema.StringAttribute{
		Description: "The username to use for HTTP basic authentication when accessing the Kubernetes API.",
		Optional:    true,
	},
	"password": schema.StringAttribute{
		Description: "The password to use for HTTP basic authentication when accessing the Kubernetes API.",
		Optional:    true,
	},
	"token": schema.StringAttribute{
		Description: "Token of your service account.",
		Optional:    true,
	},
	"insecure": schema.BoolAttribute{
		Description: "Whether the server should be accessed without verifying the TLS certificate.",
		Optional:    true,
	},
	"client_certificate": schema.StringAttribute{
		Description: "PEM-encoded client certificate for TLS authentication.",
		Optional:    true,
	},
	"client_key": schema.StringAttribute{
		Description: "PEM-encoded client certificate key for TLS authentication.",
		Optional:    true,
	},
	"cluster_ca_certificate": schema.StringAttribute{
		Description: "PEM-encoded root certificates bundle for TLS authentication.",
		Optional:    true,
	},
	"config_context": schema.StringAttribute{
		Description: "Context to choose from the config file.",
		Optional:    true,
	},
}
