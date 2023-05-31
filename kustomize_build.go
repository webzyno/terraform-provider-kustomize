package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type KustomizeBuildDataSource struct {
	kustomizer *krusty.Kustomizer
}

func NewKustomizeBuild() datasource.DataSource {
	return &KustomizeBuildDataSource{}
}

func (d *KustomizeBuildDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_build"
}

func (d *KustomizeBuildDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = kustomizeBuildSchema
}

func (d *KustomizeBuildDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	// Get kustomizer from provider data
	kustomizer, ok := req.ProviderData.(*krusty.Kustomizer)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *krusty.Kustomizer, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.kustomizer = kustomizer
}

func (d *KustomizeBuildDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Retrieve data source config
	var data KustomizeBuildModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	data.Id = types.StringValue("test")

	// Create overlay filesystem and write kustomization.yaml
	fs, err := NewOverlayFS()
	if err != nil {
		resp.Diagnostics.AddError("Failed to create OverlayFS", err.Error())
		return
	}
	kustomization := toKustomization(data)
	kustomizationContent, err := yaml.Marshal(&kustomization)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error during marshaling kustomization.yaml",
			err.Error(),
		)
	}
	if err := fs.WriteFile(KUSTOMIZATION, kustomizationContent); err != nil {
		resp.Diagnostics.AddError(
			"Failed to write kustomization.yaml to file system",
			err.Error(),
		)
	}

	resMap, err := d.kustomizer.Run(fs, ".")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error during kustomize build",
			err.Error(),
		)
	}

	manifests, err := resMap.AsYaml()
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert generated manifests to YAML", err.Error())
	}
	data.Yaml = types.StringValue(string(manifests))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
