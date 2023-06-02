package build

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
	ktypes "sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/resid"

	"github.com/webzyno/terraform-provider-kustomize/kustomize"
)

type KustomizeBuildModel struct {
	// Custom fileds
	Id   types.String `tfsdk:"id"`
	Yaml types.String `tfsdk:"yaml"`

	CommonAnnotations  map[string]string           `tfsdk:"common_annotations"`
	BuildMetadata      []types.String              `tfsdk:"build_metadata"`
	CommonLabels       map[string]string           `tfsdk:"common_labels"`
	ConfigMapGenerator []kustomize.ConfigMapArgs   `tfsdk:"config_map_generator"`
	Configurations     []types.String              `tfsdk:"configurations"`
	Crds               []types.String              `tfsdk:"crds"`
	GeneratorOptions   *kustomize.GeneratorOptions `tfsdk:"generator_options"`
	Generators         []types.String              `tfsdk:"generators"`
	HelmCharts         []kustomize.HelmChart       `tfsdk:"helm_charts"`
	HelmGlobals        *kustomize.HelmGlobals      `tfsdk:"helm_globals"`
	Images             []kustomize.Image           `tfsdk:"images"`
	Labels             []kustomize.Labels          `tfsdk:"labels"`
	NamePrefix         types.String                `tfsdk:"name_prefix"`
	NameSuffix         types.String                `tfsdk:"name_suffix"`
	Namespace          types.String                `tfsdk:"namespace"`
	Replacements       []kustomize.Replacements    `tfsdk:"replacements"`
	Openapi            map[string]string           `tfsdk:"openapi"`
	Patches            []kustomize.Patch           `tfsdk:"patches"`
	Replicas           []kustomize.Replicas        `tfsdk:"replicas"`
	Resources          []types.String              `tfsdk:"resources"`
	Components         []types.String              `tfsdk:"components"`
	SecretGenerator    []kustomize.SecretArgs      `tfsdk:"secret_generator"`
	Transformers       []types.String              `tfsdk:"transformers"`
	Validators         []types.String              `tfsdk:"validators"`
}

func ToKustomization(model KustomizeBuildModel) ktypes.Kustomization {
	return ktypes.Kustomization{
		TypeMeta: ktypes.TypeMeta{
			APIVersion: "kustomize.config.k8s.io/v1beta1",
			Kind:       "Kustomization",
		},
		OpenAPI:      model.Openapi,
		NamePrefix:   model.NamePrefix.ValueString(),
		NameSuffix:   model.NameSuffix.ValueString(),
		Namespace:    model.Namespace.ValueString(),
		CommonLabels: model.CommonLabels,
		Labels: lo.Map(model.Labels, func(label kustomize.Labels, i int) ktypes.Label {
			return ktypes.Label{
				Pairs:            label.Pairs,
				IncludeSelectors: label.IncludeSelectors.ValueBool(),
				IncludeTemplates: label.IncludeTemplates.ValueBool(),
				FieldSpecs: lo.Map(label.Fields, func(spec kustomize.FieldSpec, j int) ktypes.FieldSpec {
					return ktypes.FieldSpec{
						Gvk: resid.Gvk{
							Group:   spec.Group.ValueString(),
							Version: spec.Version.ValueString(),
							Kind:    spec.Kind.ValueString(),
						},
						Path:               spec.Path.ValueString(),
						CreateIfNotPresent: spec.Create.ValueBool(),
					}
				}),
			}
		}),
		CommonAnnotations: model.CommonAnnotations,
		Patches: lo.Map(model.Patches, func(patch kustomize.Patch, i int) ktypes.Patch {
			// Skip Options attributes since it doesn't exist in JSON schema or docs
			return ktypes.Patch{
				Path:  patch.Path.ValueString(),
				Patch: patch.Patch.ValueString(),
				Target: lo.TernaryF(patch.Target == nil, func() *ktypes.Selector { return nil }, func() *ktypes.Selector {
					return &ktypes.Selector{
						ResId: resid.ResId{
							Gvk: resid.Gvk{
								Group:   patch.Target.Group.ValueString(),
								Version: patch.Target.Version.ValueString(),
								Kind:    patch.Target.Kind.ValueString(),
							},
							Name:      patch.Target.Name.ValueString(),
							Namespace: patch.Target.Namespace.ValueString(),
						},
						AnnotationSelector: patch.Target.AnnotationSelector.ValueString(),
						LabelSelector:      patch.Target.LabelSelector.ValueString(),
					}
				}),
			}
		}),
		Images: lo.Map(model.Images, func(image kustomize.Image, i int) ktypes.Image {
			return ktypes.Image{
				Name:    image.Name.ValueString(),
				NewName: image.NewName.ValueString(),
				//TagSuffix: ",
				NewTag: image.NewTag.ValueString(),
				Digest: image.Digest.ValueString(),
			}
		}),
		Replacements: lo.Map(model.Replacements, func(replacement kustomize.Replacements, i int) ktypes.ReplacementField {
			return ktypes.ReplacementField{
				Replacement: ktypes.Replacement{
					Source: lo.TernaryF(replacement.Source == nil, func() *ktypes.SourceSelector { return nil }, func() *ktypes.SourceSelector {
						return &ktypes.SourceSelector{
							ResId: resid.ResId{
								Gvk: resid.Gvk{
									Group:   replacement.Source.Group.ValueString(),
									Version: replacement.Source.Version.ValueString(),
									Kind:    replacement.Source.Kind.ValueString(),
								},
								Name:      replacement.Source.Name.ValueString(),
								Namespace: replacement.Source.Namespace.ValueString(),
							},
							FieldPath: replacement.Source.FieldPath.ValueString(),
							Options: lo.TernaryF(replacement.Source.Options == nil, func() *ktypes.FieldOptions { return nil }, func() *ktypes.FieldOptions {
								return &ktypes.FieldOptions{
									Delimiter: replacement.Source.Options.Delimiter.ValueString(),
									Index:     int(replacement.Source.Options.Index.ValueInt64()),
									Create:    replacement.Source.Options.Create.ValueBool(),
								}
							}),
						}
					}),
					Targets: lo.Map(replacement.Targets, func(target kustomize.ReplacementsInlineTarget, j int) *ktypes.TargetSelector {
						return &ktypes.TargetSelector{
							Select: lo.TernaryF(target.Select == nil, func() *ktypes.Selector { return nil }, func() *ktypes.Selector {
								return &ktypes.Selector{
									ResId: resid.ResId{
										Gvk: resid.Gvk{
											Group:   target.Select.Group.ValueString(),
											Version: target.Select.Version.ValueString(),
											Kind:    target.Select.Kind.ValueString(),
										},
										Name:      target.Select.Name.ValueString(),
										Namespace: target.Select.Namespace.ValueString(),
									},
									//AnnotationSelector: "",
									//LabelSelector:      "",
								}
							}),
							Reject: lo.Map(target.Reject, func(reject kustomize.ReplacementsInlineTargetObject, k int) *ktypes.Selector {
								return &ktypes.Selector{
									ResId: resid.ResId{
										Gvk: resid.Gvk{
											Group:   reject.Group.ValueString(),
											Version: reject.Version.ValueString(),
											Kind:    reject.Kind.ValueString(),
										},
										Name:      reject.Name.ValueString(),
										Namespace: reject.Namespace.ValueString(),
									},
									//AnnotationSelector: "",
									//LabelSelector:      "",
								}
							}),
							FieldPaths: lo.Map(target.FieldPaths, toString),
							Options: lo.TernaryF(target.Options == nil, func() *ktypes.FieldOptions { return nil }, func() *ktypes.FieldOptions {
								return &ktypes.FieldOptions{
									Delimiter: target.Options.Delimiter.ValueString(),
									Index:     int(target.Options.Index.ValueInt64()),
									//Encoding:  "",
									Create: target.Options.Create.ValueBool(),
								}
							}),
						}
					}),
				},
				Path: replacement.Path.ValueString(),
			}
		}),
		Replicas: lo.Map(model.Replicas, func(replica kustomize.Replicas, i int) ktypes.Replica {
			return ktypes.Replica{
				Name:  replica.Name.ValueString(),
				Count: replica.Count.ValueInt64(),
			}
		}),
		Resources:  lo.Map(model.Resources, toString),
		Components: lo.Map(model.Components, toString),
		Crds:       lo.Map(model.Crds, toString),
		ConfigMapGenerator: lo.Map(model.ConfigMapGenerator, func(g kustomize.ConfigMapArgs, i int) ktypes.ConfigMapArgs {
			return ktypes.ConfigMapArgs{
				GeneratorArgs: ktypes.GeneratorArgs{
					Namespace: g.Namespace.ValueString(),
					Name:      g.Name.ValueString(),
					Behavior:  g.Behavior.ValueString(),
					KvPairSources: ktypes.KvPairSources{
						LiteralSources: lo.Map(g.Literals, toString),
						FileSources:    lo.Map(g.Files, toString),
						EnvSources:     lo.Map(g.Envs, toString),
					},
					Options: toGeneratorOptions(g.Options),
				},
			}
		}),
		SecretGenerator: lo.Map(model.SecretGenerator, func(g kustomize.SecretArgs, i int) ktypes.SecretArgs {
			return ktypes.SecretArgs{
				GeneratorArgs: ktypes.GeneratorArgs{
					Namespace: g.Namespace.ValueString(),
					Name:      g.Name.ValueString(),
					Behavior:  g.Behavior.ValueString(),
					KvPairSources: ktypes.KvPairSources{
						LiteralSources: lo.Map(g.Literals, toString),
						FileSources:    lo.Map(g.Files, toString),
						EnvSources:     lo.Map(g.Envs, toString),
					},
					Options: toGeneratorOptions(g.Options),
				},
				Type: g.Type.ValueString(),
			}
		}),
		HelmGlobals: lo.TernaryF(model.HelmGlobals == nil, func() *ktypes.HelmGlobals { return nil }, func() *ktypes.HelmGlobals {
			return &ktypes.HelmGlobals{
				ChartHome:  model.HelmGlobals.ChartHome.ValueString(),
				ConfigHome: model.HelmGlobals.ConfigHome.ValueString(),
			}
		}),
		HelmCharts: lo.Map(model.HelmCharts, func(chart kustomize.HelmChart, i int) ktypes.HelmChart {
			return ktypes.HelmChart{
				Name:                  chart.Name.ValueString(),
				Version:               chart.Version.ValueString(),
				Repo:                  chart.Repo.ValueString(),
				ReleaseName:           chart.ReleaseName.ValueString(),
				Namespace:             chart.Namespace.ValueString(),
				AdditionalValuesFiles: lo.Map(chart.AdditionalValuesFiles, toString),
				ValuesFile:            chart.ValuesFile.ValueString(),
				ValuesInline:          chart.ValuesInline,
				ValuesMerge:           chart.ValuesMerge.ValueString(),
				IncludeCRDs:           chart.IncludeCRDs.ValueBool(),
				ApiVersions:           lo.Map(chart.ApiVersions, toString),
				NameTemplate:          chart.NameTemplate.ValueString(),
				SkipTests:             chart.SkipTests.ValueBool(),
			}
		}),
		GeneratorOptions: toGeneratorOptions(model.GeneratorOptions),
		Configurations:   lo.Map(model.Configurations, toString),
		Generators:       lo.Map(model.Generators, toString),
		Transformers:     lo.Map(model.Transformers, toString),
		Validators:       lo.Map(model.Validators, toString),
		BuildMetadata:    lo.Map(model.BuildMetadata, toString),
	}
}

func toString(tStr types.String, _ int) string {
	return tStr.ValueString()
}

func toGeneratorOptions(options *kustomize.GeneratorOptions) *ktypes.GeneratorOptions {
	if options == nil {
		return nil
	}
	return &ktypes.GeneratorOptions{
		Annotations:           options.Annotations,
		DisableNameSuffixHash: options.DisableNameSuffixHash.ValueBool(),
		Immutable:             options.Immutable.ValueBool(),
		Labels:                options.Labels,
	}
}
