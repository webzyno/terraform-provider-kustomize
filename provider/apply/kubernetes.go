package apply

import (
	"errors"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/kustomize/api/resource"
)

func checkAPIResourceIsPresent(supportedResources []*metav1.APIResourceList, resource *resource.Resource) (*metav1.APIResource, bool) {
	for _, srList := range supportedResources {
		if srList == nil {
			continue
		}
		rGvk := resource.GetGvk()
		if rGvk.ApiVersion() == srList.GroupVersion {
			for _, sr := range srList.APIResources {
				if rGvk.Kind == sr.Kind {
					sr.Group = srList.GroupVersion
					sr.Kind = srList.Kind
					return &sr, true
				}
			}
		}
	}
	return nil, false
}

func getRestClientFromResource(dynamicClient dynamic.Interface, supportedResources []*metav1.APIResourceList, resource *resource.Resource) (dynamic.ResourceInterface, error) {
	apiResource, exists := checkAPIResourceIsPresent(supportedResources, resource)
	if !exists {
		return nil, errors.New("the Kubernetes API server doesn't support this resource")
	}

	resourceSchema := schema.GroupVersionResource{Group: apiResource.Group, Version: apiResource.Version, Resource: apiResource.Name}
	// For core services (ServiceAccount, Service etc) the group is incorrectly parsed.
	// "v1" should be empty group and "v1" for version
	if resourceSchema.Group == "v1" && resourceSchema.Version == "" {
		resourceSchema.Group = ""
		resourceSchema.Version = "v1"
	}

	restClient := dynamicClient.Resource(resourceSchema)

	if apiResource.Namespaced {
		if resource.GetNamespace() == "" {
			return restClient.Namespace("default"), nil
		}
		return restClient.Namespace(resource.GetNamespace()), nil
	}

	return restClient, nil
}

func kustomizeResourceToUnstructured(resource *resource.Resource) (*unstructured.Unstructured, error) {
	json, err := resource.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to convert resource to JSON: %w", err)
	}
	us := &unstructured.Unstructured{}
	if err := us.UnmarshalJSON(json); err != nil {
		return nil, fmt.Errorf("Failed to convert JSON to *unstructured.Unstructured: %w", err)
	}
	return us, nil
}
