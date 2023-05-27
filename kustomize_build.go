package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type KustomizeBuildDataSource struct {
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

func (d *KustomizeBuildDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data KustomizeBuildModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	data.Id = types.StringValue("test")

	tflog.Debug(ctx, "common_annotations")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
