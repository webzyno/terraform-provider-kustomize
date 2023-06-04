package apply

import (
	"fmt"

	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/kyaml/yaml"

	"github.com/webzyno/terraform-provider-kustomize/virtfs"
)

func kustomizeBuild(kustomizer *krusty.Kustomizer, model KustomizeApplyModel) (resmap.ResMap, error) {
	// Create overlay filesystem and write kustomization.yaml
	fs, err := virtfs.NewOverlayFS()
	if err != nil {
		return nil, fmt.Errorf("failed to create OverlayFS: %w", err)
	}
	kustomization := ToKustomization(model)
	kustomizationContent, err := yaml.Marshal(&kustomization)
	if err != nil {
		return nil, fmt.Errorf("unexpected error during marshaling kustomization.yaml: %w", err)
	}
	if err := fs.WriteFile(virtfs.KUSTOMIZATION, kustomizationContent); err != nil {
		return nil, fmt.Errorf("failed to write kustomization.yaml to file system: %w", err)
	}

	return kustomizer.Run(fs, ".")
}
