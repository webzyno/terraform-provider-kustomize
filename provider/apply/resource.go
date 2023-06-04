package apply

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/kustomize/api/krusty"

	"github.com/webzyno/terraform-provider-kustomize/client"
)

const FieldManager = "terraform-provider-kustomize"

type KustomizeApplyResource struct {
	kustomizer      *krusty.Kustomizer
	dynamicClient   dynamic.Interface
	discoveryClient discovery.DiscoveryInterface
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
	clientSet, ok := req.ProviderData.(*client.ClientSet)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *krusty.Kustomizer, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.kustomizer = clientSet.Kustomizer
	r.dynamicClient = clientSet.DynamicClient
	r.discoveryClient = clientSet.DiscoveryClient
}

func (r *KustomizeApplyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve resource from plan data
	var data KustomizeApplyModel
	if resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}
	// Set id for test purpose
	data.Id = types.StringValue("test")

	// Run kustomize build and save resmap in yaml
	resMap, err := kustomizeBuild(r.kustomizer, data)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to run kustomize build",
			err.Error(),
		)
		return
	}
	manifests, err := resMap.AsYaml()
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert generated manifests to YAML", err.Error())
		return
	}
	data.Yaml = types.StringValue(string(manifests))

	// Get server supported resources
	_, resources, err := r.discoveryClient.ServerGroupsAndResources()
	if err != nil && !discovery.IsGroupDiscoveryFailedError(err) {
		resp.Diagnostics.AddError("Failed to get server groups and resources", err.Error())
		return
	}

	var objects []string
	for _, res := range resMap.Resources() {
		restClient, err := getRestClientFromResource(r.dynamicClient, resources, res)
		if err != nil {
			resp.Diagnostics.AddError("Failed to create rest client from resource", err.Error())
			return
		}

		// Convert resource to unstructured
		unstructured, err := kustomizeResourceToUnstructured(res)
		if err != nil {
			resp.Diagnostics.AddError("Failed to convert Kustomize resource to unstructured", err.Error())
			return
		}

		createdUnstructured, err := restClient.Apply(ctx, unstructured.GetName(), unstructured, metav1.ApplyOptions{FieldManager: FieldManager})
		if err != nil {
			resp.Diagnostics.AddError("Failed to create object", err.Error())
			return
		}
		createdUsJson, err := createdUnstructured.MarshalJSON()
		if err != nil {
			resp.Diagnostics.AddError("failed to convert created object to JSON", err.Error())
			return
		}

		objects = append(objects, string(createdUsJson))
	}

	// Set objects to model
	objectsModel, diagnostics := types.ListValueFrom(ctx, types.StringType, &objects)
	data.Objects = objectsModel
	if resp.Diagnostics.Append(diagnostics...); resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.Append(resp.State.Set(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}
}

func (r *KustomizeApplyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data KustomizeApplyModel
	if resp.Diagnostics.Append(req.State.Get(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}

	// Run kustomize build and save resmap in yaml
	resMap, err := kustomizeBuild(r.kustomizer, data)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to run kustomize build",
			err.Error(),
		)
		return
	}
	manifests, err := resMap.AsYaml()
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert generated manifests to YAML", err.Error())
		return
	}
	data.Yaml = types.StringValue(string(manifests))

	// Get server supported resources
	_, resources, err := r.discoveryClient.ServerGroupsAndResources()
	if err != nil && !discovery.IsGroupDiscoveryFailedError(err) {
		resp.Diagnostics.AddError("Failed to get server groups and resources", err.Error())
		return
	}

	var objects []string
	for _, res := range resMap.Resources() {
		restClient, err := getRestClientFromResource(r.dynamicClient, resources, res)
		if err != nil {
			resp.Diagnostics.AddError("Failed to create rest client from resource", err.Error())
			return
		}

		// Convert resource to unstructured
		unstructured, err := kustomizeResourceToUnstructured(res)
		if err != nil {
			resp.Diagnostics.AddError("Failed to convert Kustomize resource to unstructured", err.Error())
			return
		}

		getUnstructured, err := restClient.Get(ctx, unstructured.GetName(), metav1.GetOptions{})
		if err != nil {
			resp.Diagnostics.AddError("Failed to get object", err.Error())
			return
		}
		getUnstructuredJson, err := getUnstructured.MarshalJSON()
		if err != nil {
			resp.Diagnostics.AddError("failed to convert get object to JSON", err.Error())
			return
		}

		objects = append(objects, string(getUnstructuredJson))
	}

	// Set objects to model
	objectsModel, diagnostics := types.ListValueFrom(ctx, types.StringType, &objects)
	data.Objects = objectsModel
	if resp.Diagnostics.Append(diagnostics...); resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.Append(resp.State.Set(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}
}

func (r *KustomizeApplyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve resource from plan data
	var data KustomizeApplyModel
	if resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}

	// Run kustomize build and save resmap in yaml
	resMap, err := kustomizeBuild(r.kustomizer, data)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to run kustomize build",
			err.Error(),
		)
		return
	}
	manifests, err := resMap.AsYaml()
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert generated manifests to YAML", err.Error())
		return
	}
	data.Yaml = types.StringValue(string(manifests))

	// Get server supported resources
	_, resources, err := r.discoveryClient.ServerGroupsAndResources()
	if err != nil && !discovery.IsGroupDiscoveryFailedError(err) {
		resp.Diagnostics.AddError("Failed to get server groups and resources", err.Error())
		return
	}

	var objects []string
	for _, res := range resMap.Resources() {
		restClient, err := getRestClientFromResource(r.dynamicClient, resources, res)
		if err != nil {
			resp.Diagnostics.AddError("Failed to create rest client from resource", err.Error())
			return
		}

		// Convert resource to unstructured
		unstructured, err := kustomizeResourceToUnstructured(res)
		if err != nil {
			resp.Diagnostics.AddError("Failed to convert Kustomize resource to unstructured", err.Error())
			return
		}

		createdUnstructured, err := restClient.Apply(ctx, unstructured.GetName(), unstructured, metav1.ApplyOptions{FieldManager: FieldManager})
		if err != nil {
			resp.Diagnostics.AddError("Failed to update object", err.Error())
			return
		}
		appliedJson, err := createdUnstructured.MarshalJSON()
		if err != nil {
			resp.Diagnostics.AddError("failed to convert created object to JSON", err.Error())
			return
		}

		objects = append(objects, string(appliedJson))
	}

	// Set objects to model
	objectsModel, diagnostics := types.ListValueFrom(ctx, types.StringType, &objects)
	data.Objects = objectsModel
	if resp.Diagnostics.Append(diagnostics...); resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.Append(resp.State.Set(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}
}

func (r *KustomizeApplyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve resource from state data
	var data KustomizeApplyModel
	if resp.Diagnostics.Append(req.State.Get(ctx, &data)...); resp.Diagnostics.HasError() {
		return
	}

	// Run kustomize build
	resMap, err := kustomizeBuild(r.kustomizer, data)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to run kustomize build",
			err.Error(),
		)
		return
	}

	// Get server supported resources
	_, resources, err := r.discoveryClient.ServerGroupsAndResources()
	if err != nil && !discovery.IsGroupDiscoveryFailedError(err) {
		resp.Diagnostics.AddError("Failed to get server groups and resources", err.Error())
		return
	}

	for _, res := range lo.Reverse(resMap.Resources()) {
		restClient, err := getRestClientFromResource(r.dynamicClient, resources, res)
		if err != nil {
			resp.Diagnostics.AddError("Failed to create rest client from resource", err.Error())
			return
		}

		if err := restClient.Delete(ctx, res.GetName(), metav1.DeleteOptions{}); err != nil {
			resp.Diagnostics.AddError("Failed to delete Kubernetes object", err.Error())
			return
		}
	}
}
