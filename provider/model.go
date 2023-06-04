package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type KustomizeProviderModel struct {
	Kubernetes *KubernetesModel `tfsdk:"kubernetes"`
}

type KubernetesModel struct {
	ConfigPath           types.String   `tfsdk:"config_path"`
	ConfigPaths          []types.String `tfsdk:"config_paths"`
	Host                 types.String   `tfsdk:"host"`
	Username             types.String   `tfsdk:"username"`
	Password             types.String   `tfsdk:"password"`
	Token                types.String   `tfsdk:"token"`
	Insecure             types.Bool     `tfsdk:"insecure"`
	ClientCertificate    types.String   `tfsdk:"client_certificate"`
	ClientKey            types.String   `tfsdk:"client_key"`
	ClusterCACertificate types.String   `tfsdk:"cluster_ca_certificate"`
	ConfigContext        types.String   `tfsdk:"config_context"`
}
