package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"sigs.k8s.io/kustomize/api/krusty"
)

func NewProvider() provider.Provider {
	return &KustomizeProvider{}
}

type KustomizeProvider struct {
	kustomizer *krusty.Kustomizer
}

func (p *KustomizeProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "kustomize"
}

func (p *KustomizeProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `This Kustomization provider is used to build Kubernetes manifests using Kustomization.
Although there are existing providers, this provider gives you the best DX and mitigate the datasource's read when apply issue.`,
	}
}

func (p *KustomizeProvider) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Create Kustomize client
	p.kustomizer = krusty.MakeKustomizer(krusty.MakeDefaultOptions())

	// Make kustomizer available to data source
	resp.DataSourceData = p.kustomizer
}

func (p *KustomizeProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewKustomizeBuild,
	}
}

func (p *KustomizeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
