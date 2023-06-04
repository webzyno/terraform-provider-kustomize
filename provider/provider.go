package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"sigs.k8s.io/kustomize/api/krusty"

	"github.com/webzyno/terraform-provider-kustomize/client"
	"github.com/webzyno/terraform-provider-kustomize/provider/apply"
	"github.com/webzyno/terraform-provider-kustomize/provider/build"
)

var TestAccProtoV6Providers = map[string]func() (tfprotov6.ProviderServer, error){
	"kustomize": providerserver.NewProtocol6WithError(NewProvider()),
}

func NewProvider() provider.Provider {
	return &KustomizeProvider{}
}

type KustomizeProvider struct {
	clientSet *client.ClientSet
}

func (p *KustomizeProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "kustomize"
}

func (p *KustomizeProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = kustomizeProviderSchema
}

func (p *KustomizeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data KustomizeProviderModel
	if resp.Diagnostics.Append(req.Config.Get(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}

	clientSet := &client.ClientSet{}

	// Create Kustomize client
	kustomizeOptions := krusty.MakeDefaultOptions()
	kustomizeOptions.Reorder = krusty.ReorderOptionLegacy
	clientSet.Kustomizer = krusty.MakeKustomizer(kustomizeOptions)

	// Create Kubernetes dynamic client if kubernetes is set
	if data.Kubernetes != nil {
		dynamicClient, err := newDynamicClient(*data.Kubernetes)
		if err != nil {
			resp.Diagnostics.AddError("Failed to create Kubernetes client.", err.Error())
			return
		}
		clientSet.DynamicClient = dynamicClient

		discoveryClient, err := newDiscoveryClient(*data.Kubernetes)
		if err != nil {
			resp.Diagnostics.AddError("Failed to create Kubernetes discovery client.", err.Error())
			return
		}
		clientSet.DiscoveryClient = discoveryClient
	}

	p.clientSet = clientSet

	// Make kustomizer available to data source and resource
	resp.DataSourceData = p.clientSet
	resp.ResourceData = p.clientSet
}

func (p *KustomizeProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		build.NewKustomizeBuild,
	}
}

func (p *KustomizeProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		apply.NewKustomizeApply,
	}
}
