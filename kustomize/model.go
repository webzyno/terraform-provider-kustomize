package kustomize

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
	ktypes "sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/resid"
)

// Reference: Kustomize JSON schema https://github.com/SchemaStore/schemastore/blob/master/src/schemas/json/kustomization.json
// Deprecated attributes are removed
// Skip Inventory attributes because we can't find any documentation in Kustomize
// Skip Kind and Metadata attributes because they are Kubernetes required attributes and have no effect

type Model struct {
	CommonAnnotations  map[string]string `tfsdk:"common_annotations"`
	BuildMetadata      []types.String    `tfsdk:"build_metadata"`
	CommonLabels       map[string]string `tfsdk:"common_labels"`
	ConfigMapGenerator []ConfigMapArgs   `tfsdk:"config_map_generator"`
	Configurations     []types.String    `tfsdk:"configurations"`
	Crds               []types.String    `tfsdk:"crds"`
	GeneratorOptions   *GeneratorOptions `tfsdk:"generator_options"`
	Generators         []types.String    `tfsdk:"generators"`
	HelmCharts         []HelmChart       `tfsdk:"helm_charts"`
	HelmGlobals        *HelmGlobals      `tfsdk:"helm_globals"`
	Images             []Image           `tfsdk:"images"`
	Labels             []Labels          `tfsdk:"labels"`
	NamePrefix         types.String      `tfsdk:"name_prefix"`
	NameSuffix         types.String      `tfsdk:"name_suffix"`
	Namespace          types.String      `tfsdk:"namespace"`
	Replacements       []Replacements    `tfsdk:"replacements"`
	Openapi            map[string]string `tfsdk:"openapi"`
	Patches            []Patch           `tfsdk:"patches"`
	Replicas           []Replicas        `tfsdk:"replicas"`
	Resources          []types.String    `tfsdk:"resources"`
	Components         []types.String    `tfsdk:"components"`
	SecretGenerator    []SecretArgs      `tfsdk:"secret_generator"`
	Transformers       []types.String    `tfsdk:"transformers"`
	Validators         []types.String    `tfsdk:"validators"`
}

type ConfigMapArgs struct {
	Behavior  types.String      `tfsdk:"behavior"`
	Envs      []types.String    `tfsdk:"envs"`
	Files     []types.String    `tfsdk:"files"`
	Literals  []types.String    `tfsdk:"literals"`
	Name      types.String      `tfsdk:"name"`
	Namespace types.String      `tfsdk:"namespace"`
	Options   *GeneratorOptions `tfsdk:"options"`
}

type GeneratorOptions struct {
	Annotations           map[string]string `tfsdk:"annotations"`
	DisableNameSuffixHash types.Bool        `tfsdk:"disable_name_suffix_hash"`
	Immutable             types.Bool        `tfsdk:"immutable"`
	Labels                map[string]string `tfsdk:"labels"`
}

type HelmChart struct {
	Name                  types.String   `tfsdk:"name"`
	Version               types.String   `tfsdk:"version"`
	Repo                  types.String   `tfsdk:"repo"`
	ReleaseName           types.String   `tfsdk:"release_name"`
	Namespace             types.String   `tfsdk:"namespace"`
	ValuesFile            types.String   `tfsdk:"values_file"`
	ValuesInline          map[string]any `tfsdk:"values_inline"` // The JSON schema is an arbitrary object, but there is no corresponding type
	ValuesMerge           types.String   `tfsdk:"values_merge"`
	IncludeCRDs           types.Bool     `tfsdk:"include_crds"`
	AdditionalValuesFiles []types.String `tfsdk:"additional_values_files"`
	SkipTests             types.Bool     `tfsdk:"skip_tests"`
	ApiVersions           []types.String `tfsdk:"api_versions"`
	NameTemplate          types.String   `tfsdk:"name_template"`
}

type HelmGlobals struct {
	ChartHome  types.String `tfsdk:"chart_home"`
	ConfigHome types.String `tfsdk:"config_home"`
}

type Image struct {
	Digest  types.String `tfsdk:"digest"`
	Name    types.String `tfsdk:"name"`
	NewName types.String `tfsdk:"new_name"`
	NewTag  types.String `tfsdk:"new_tag"`
}

type Labels struct {
	Pairs            map[string]string `tfsdk:"pairs"`
	IncludeSelectors types.Bool        `tfsdk:"include_selectors"`
	IncludeTemplates types.Bool        `tfsdk:"include_templates"`
	Fields           []FieldSpec       `tfsdk:"fields"`
}

type FieldSpec struct {
	Create  types.Bool   `tfsdk:"create"`
	Group   types.String `tfsdk:"group"`
	Kind    types.String `tfsdk:"kind"`
	Path    types.String `tfsdk:"path"`
	Version types.String `tfsdk:"version"`
}

type Replacements struct {
	Path    types.String               `tfsdk:"path"`
	Source  *ReplacementsInlineSource  `tfsdk:"source"`
	Targets []ReplacementsInlineTarget `tfsdk:"targets"`
}

type ReplacementsInlineSource struct {
	Group     types.String         `tfsdk:"group"`
	Version   types.String         `tfsdk:"version"`
	Kind      types.String         `tfsdk:"kind"`
	Name      types.String         `tfsdk:"name"`
	Namespace types.String         `tfsdk:"namespace"`
	FieldPath types.String         `tfsdk:"field_path"`
	Options   *ReplacementsOptions `tfsdk:"options"`
}

type ReplacementsInlineTarget struct {
	Select     *ReplacementsInlineTargetObject  `tfsdk:"select"`
	Reject     []ReplacementsInlineTargetObject `tfsdk:"reject"`
	FieldPaths []types.String                   `tfsdk:"field_paths"`
	Options    *ReplacementsOptions             `tfsdk:"options"`
}

type ReplacementsOptions struct {
	Delimiter types.String `tfsdk:"delimiter"`
	Index     types.Int64  `tfsdk:"index"`
	Create    types.Bool   `tfsdk:"create"`
}

type ReplacementsInlineTargetObject struct {
	Group     types.String `tfsdk:"group"`
	Version   types.String `tfsdk:"version"`
	Kind      types.String `tfsdk:"kind"`
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

type Patch struct {
	Path   types.String         `tfsdk:"path"`
	Patch  types.String         `tfsdk:"patch"`
	Target *PatchTargetOptional `tfsdk:"target"`
}

type PatchTargetOptional struct {
	Group              types.String `tfsdk:"group"`
	Kind               types.String `tfsdk:"kind"`
	Name               types.String `tfsdk:"name"`
	Namespace          types.String `tfsdk:"namespace"`
	Version            types.String `tfsdk:"version"`
	LabelSelector      types.String `tfsdk:"label_selector"`
	AnnotationSelector types.String `tfsdk:"annotation_selector"`
}

type Replicas struct {
	Name  types.String `tfsdk:"name"`
	Count types.Int64  `tfsdk:"count"`
}

type SecretArgs struct {
	Behavior  types.String      `tfsdk:"behavior"`
	Env       types.String      `tfsdk:"env"`
	Envs      []types.String    `tfsdk:"envs"`
	Files     []types.String    `tfsdk:"files"`
	Literals  []types.String    `tfsdk:"literals"`
	Name      types.String      `tfsdk:"name"`
	Namespace types.String      `tfsdk:"namespace"`
	Options   *GeneratorOptions `tfsdk:"options"`
	Type      types.String      `tfsdk:"type"`
}

func ToKustomization(model Model) ktypes.Kustomization {
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
		Labels: lo.Map(model.Labels, func(label Labels, i int) ktypes.Label {
			return ktypes.Label{
				Pairs:            label.Pairs,
				IncludeSelectors: label.IncludeSelectors.ValueBool(),
				IncludeTemplates: label.IncludeTemplates.ValueBool(),
				FieldSpecs: lo.Map(label.Fields, func(spec FieldSpec, j int) ktypes.FieldSpec {
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
		Patches: lo.Map(model.Patches, func(patch Patch, i int) ktypes.Patch {
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
		Images: lo.Map(model.Images, func(image Image, i int) ktypes.Image {
			return ktypes.Image{
				Name:    image.Name.ValueString(),
				NewName: image.NewName.ValueString(),
				//TagSuffix: ",
				NewTag: image.NewTag.ValueString(),
				Digest: image.Digest.ValueString(),
			}
		}),
		Replacements: lo.Map(model.Replacements, func(replacement Replacements, i int) ktypes.ReplacementField {
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
					Targets: lo.Map(replacement.Targets, func(target ReplacementsInlineTarget, j int) *ktypes.TargetSelector {
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
							Reject: lo.Map(target.Reject, func(reject ReplacementsInlineTargetObject, k int) *ktypes.Selector {
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
		Replicas: lo.Map(model.Replicas, func(replica Replicas, i int) ktypes.Replica {
			return ktypes.Replica{
				Name:  replica.Name.ValueString(),
				Count: replica.Count.ValueInt64(),
			}
		}),
		Resources:  lo.Map(model.Resources, toString),
		Components: lo.Map(model.Components, toString),
		Crds:       lo.Map(model.Crds, toString),
		ConfigMapGenerator: lo.Map(model.ConfigMapGenerator, func(g ConfigMapArgs, i int) ktypes.ConfigMapArgs {
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
		SecretGenerator: lo.Map(model.SecretGenerator, func(g SecretArgs, i int) ktypes.SecretArgs {
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
		HelmCharts: lo.Map(model.HelmCharts, func(chart HelmChart, i int) ktypes.HelmChart {
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

func toGeneratorOptions(options *GeneratorOptions) *ktypes.GeneratorOptions {
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
