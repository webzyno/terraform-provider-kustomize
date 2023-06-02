package apply

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"sigs.k8s.io/kustomize/api/krusty"
)

type KustomizeApplyResource struct {
	kustomizer *krusty.Kustomizer
}

func NewKustomizeApply() resource.Resource {
	return &KustomizeApplyResource{}
}

func (r *KustomizeApplyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apply"
}

func (r *KustomizeApplyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = kustomizeApplySchema
}

func (r *KustomizeApplyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.kustomizer = kustomizer
}

func (r *KustomizeApplyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *KustomizeApplyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *KustomizeApplyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *KustomizeApplyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//TODO implement me
	panic("implement me")
}
