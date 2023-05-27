package main

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Reference: Kustomize JSON schema https://github.com/SchemaStore/schemastore/blob/master/src/schemas/json/kustomization.json
// Deprecated attributes are removed
// Skip Inventory attributes because we can't find any documentation in Kustomize
// Skip Kind and Metadata attributes because they are Kubernetes required attributes and have no effect

type KustomizeBuildModel struct {
	Id                 types.String      `tfsdk:"id"`
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
	Name                  types.String      `tfsdk:"name"`
	Version               types.String      `tfsdk:"version"`
	Repo                  types.String      `tfsdk:"repo"`
	ReleaseName           types.String      `tfsdk:"release_name"`
	Namespace             types.String      `tfsdk:"namespace"`
	ValuesFile            types.String      `tfsdk:"values_file"`
	ValuesInline          map[string]string `tfsdk:"values_inline"` // The JSON schema is an arbitrary object, but there is no corresponding type
	ValuesMerge           types.String      `tfsdk:"values_merge"`
	IncludeCRDs           types.Bool        `tfsdk:"include_crds"`
	AdditionalValuesFiles []types.String    `tfsdk:"additional_values_files"`
	SkipTests             types.Bool        `tfsdk:"skip_tests"`
	ApiVersions           []types.String    `tfsdk:"api_versions"`
	NameTemplate          types.String      `tfsdk:"name_template"`
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

var kustomizeBuildSchema = schema.Schema{
	Description: "",
	Attributes: map[string]schema.Attribute{
		// Required for testing framework
		"id": schema.StringAttribute{
			Computed: true,
		},
		"common_annotations": schema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "common_annotations to add to all objects",
		},
		"build_metadata": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "build_metadata is a list of strings used to toggle different build options",
		},
		"common_labels": schema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "common_labels to add to all objects and selectors",
		},
		"config_map_generator": schema.ListNestedAttribute{
			Description: "config_map_generator is a list of configmaps to generate from local data (one configMap per list item)",
			Optional:    true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"behavior": schema.StringAttribute{
						Optional:    true,
						Description: "behavior configures the strategy for overriding ConfigMap",
					},
					"envs": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "A list of file paths. The contents of each file should be one key=value pair per line",
					},
					"files": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "A list of file sources to use in creating a list of key, value pairs",
					},
					"literals": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "A list of literal pair sources. Each literal source should be a key and literal value, e.g. `key=value`",
					},
					"name": schema.StringAttribute{
						Optional:    true,
						Description: "name - actually the partial name - of the generated resource",
					},
					"namespace": schema.StringAttribute{
						Optional:    true,
						Description: "namespace for the configmap, optional",
					},
					"options": generatorOptionsAttributes,
				},
			},
		},
		"configurations": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "configurations is a list of transformer configuration files",
		},
		"crds": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "crds specifies relative paths to Custom Resource Definition files. This allows custom resources to be recognized as operands, making it possible to add them to the Resources list. CRDs themselves are not modified.",
		},
		"generator_options": generatorOptionsAttributes,
		"generators": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "generators is a list of files containing custom generators",
		},
		"helm_charts": schema.ListNestedAttribute{
			Optional:    true,
			Description: "helm_charts is a list of helm chart configuration instances",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Optional: true,
					},
					"version": schema.StringAttribute{
						Optional: true,
					},
					"repo": schema.StringAttribute{
						Optional: true,
					},
					"release_name": schema.StringAttribute{
						Optional: true,
					},
					"namespace": schema.StringAttribute{
						Optional: true,
					},
					"values_file": schema.StringAttribute{
						Optional: true,
					},
					"values_inline": schema.MapAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"values_merge": schema.StringAttribute{
						Optional: true,
					},
					"include_crds": schema.BoolAttribute{
						Optional: true,
					},
					"additional_values_files": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"skip_tests": schema.BoolAttribute{
						Optional: true,
					},
					"api_versions": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"name_template": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
		"helm_globals": schema.SingleNestedAttribute{
			Optional:    true,
			Description: "helm_globals contains helm configuration that isn't chart specific",
			Attributes: map[string]schema.Attribute{
				"chart_home": schema.StringAttribute{
					Optional:    true,
					Description: "chart_home is a file path, relative to the kustomization root, to a directory containing a subdirectory for each chart to be included in the kustomization",
				},
				"config_home": schema.StringAttribute{
					Optional:    true,
					Description: "config_home defines a value that kustomize should pass to helm via the HELM_CONFIG_HOME environment variable",
				},
			},
		},
		"images": schema.ListNestedAttribute{
			Optional:    true,
			Description: "images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"digest": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Optional: true,
					},
					"new_name": schema.StringAttribute{
						Optional: true,
					},
					"new_tag": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
		"labels": schema.ListNestedAttribute{
			Optional:    true,
			Description: "labels to add to all objects but not selectors",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"pairs": schema.MapAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "pairs contains the key-value pairs for labels to add",
					},
					"include_selectors": schema.BoolAttribute{
						Optional:    true,
						Description: "include_selectors inidicates should transformer include the fieldSpecs for selectors",
					},
					"include_templates": schema.BoolAttribute{
						Optional:    true,
						Description: "include_templates inidicates should transformer include the template labels",
					},
					"fields": schema.ListNestedAttribute{
						Optional:    true,
						Description: "fields completely specifies a kustomizable field in a k8s API object. It helps define the operands of transformations",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"create": schema.BoolAttribute{
									Optional: true,
								},
								"group": schema.StringAttribute{
									Optional: true,
								},
								"kind": schema.StringAttribute{
									Optional: true,
								},
								"path": schema.StringAttribute{
									Optional: true,
								},
								"version": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"name_prefix": schema.StringAttribute{
			Optional:    true,
			Description: "name_prefix will prefix the names of all resources mentioned in the kustomization file including generated configmaps and secrets",
		},
		"name_suffix": schema.StringAttribute{
			Optional:    true,
			Description: "name_suffix will suffix the names of all resources mentioned in the kustomization file including generated configmaps and secrets",
		},
		"namespace": schema.StringAttribute{
			Optional:    true,
			Description: "namespace to add to all objects",
		},
		"replacements": schema.ListNestedAttribute{
			Optional:    true,
			Description: "replacements substitute field(s) in N target(s) with a field from a source",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"path": schema.StringAttribute{
						Optional: true,
					},
					"source": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "The source of the value",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Optional:    true,
								Description: "The group of the referent",
							},
							"version": schema.StringAttribute{
								Optional:    true,
								Description: "The version of the referent",
							},
							"kind": schema.StringAttribute{
								Optional:    true,
								Description: "The kind of the referent",
							},
							"name": schema.StringAttribute{
								Optional:    true,
								Description: "The name of the referent",
							},
							"namespace": schema.StringAttribute{
								Optional:    true,
								Description: "The namespace of the referent",
							},
							"field_path": schema.StringAttribute{
								Optional:    true,
								Description: "The structured path to the source value",
							},
							"options": replacementsOptionsAttributes,
						},
					},
					"targets": schema.ListNestedAttribute{
						Optional:    true,
						Description: "The N fields to write the value to",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"select": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Include objects that match this",
									Attributes:  replacementsInlineTargetObjectAttributes,
								},
								"reject": schema.ListNestedAttribute{
									Optional:    true,
									Description: "Exclude objects that match this",
									NestedObject: schema.NestedAttributeObject{
										Attributes: replacementsInlineTargetObjectAttributes,
									},
								},
								"field_paths": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Description: "The structured path(s) to the target nodes",
								},
								"options": replacementsOptionsAttributes,
							},
						},
					},
				},
			},
		},
		"openapi": schema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "openapi contains information about what kubernetes schema to use",
		},
		"patches": schema.ListNestedAttribute{
			Optional:    true,
			Description: "Apply a patch to multiple resources",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"path": schema.StringAttribute{
						Optional: true,
					},
					"patch": schema.StringAttribute{
						Optional: true,
					},
					"target": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Refers to a Kubernetes object that the patch will be applied to. It must refer to a Kubernetes resource under the purview of this kustomization",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Optional: true,
							},
							"kind": schema.StringAttribute{
								Optional: true,
							},
							"name": schema.StringAttribute{
								Optional: true,
							},
							"namespace": schema.StringAttribute{
								Optional: true,
							},
							"version": schema.StringAttribute{
								Optional: true,
							},
							"label_selector": schema.StringAttribute{
								Optional: true,
							},
							"annotation_selector": schema.StringAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"replicas": schema.ListNestedAttribute{
			Optional:    true,
			Description: "replicas is a list of (resource name, count) for changing number of replicas for a resources. It will match any group and kind that has a matching name and that is one of: Deployment, ReplicationController, Replicaset, Statefulset.",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Optional: true,
					},
					"count": schema.Int64Attribute{
						Optional: true,
					},
				},
			},
		},
		"resources": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "resources specifies relative paths to files holding YAML representations of kubernetes API objects. URLs and globs not supported.",
		},
		"components": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "components are relative paths or git repository URLs specifying a directory containing a kustomization.yaml file of Kind Component.",
		},
		"secret_generator": schema.ListNestedAttribute{
			Optional:    true,
			Description: "secret_generator is a list of secrets to generate from local data (one secret per list item)",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"behavior": schema.StringAttribute{
						Optional:    true,
						Description: "behavior configures the strategy for overriding Secret",
					},
					"env": schema.StringAttribute{
						Optional: true,
					},
					"envs": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "A list of file paths. The contents of each file should be one key=value pair per line",
					},
					"files": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "A list of file sources to use in creating a list of key, value pairs",
					},
					"literals": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "A list of literal pair sources. Each literal source should be a key and literal value, e.g. `key=value`",
					},
					"name": schema.StringAttribute{
						Optional:    true,
						Description: "name - actually the partial name - of the generated resource",
					},
					"namespace": schema.StringAttribute{
						Optional:    true,
						Description: "namespace for the secret, optional",
					},
					"options": generatorOptionsAttributes,
					"type": schema.StringAttribute{
						Optional:    true,
						Description: "type of the secret, optional",
					},
				},
			},
		},
		"transformers": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "transformers is a list of files containing transformers",
		},
		"validators": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "validators is a list of files containing validators",
		},
	},
}

var generatorOptionsAttributes = schema.SingleNestedAttribute{
	Optional:    true,
	Description: "generator_options modify behavior of all ConfigMap and Secret generators",
	Attributes: map[string]schema.Attribute{
		"annotations": schema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "annotations to add to all generated resources",
		},
		"disable_name_suffix_hash": schema.BoolAttribute{
			Optional:    true,
			Description: "disable_name_suffix_hash if true disables the default behavior of adding a suffix to the names of generated resources that is a hash of the resource contents",
		},
		"immutable": schema.BoolAttribute{
			Optional:    true,
			Description: "immutable if true add to all generated resources",
		},
		"labels": schema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "labels to add to all generated resources",
		},
	},
}

var replacementsOptionsAttributes = schema.SingleNestedAttribute{
	Optional: true,
	Attributes: map[string]schema.Attribute{
		"delimiter": schema.StringAttribute{
			Optional: true,
		},
		"index": schema.Int64Attribute{
			Optional: true,
		},
		"create": schema.BoolAttribute{
			Optional: true,
		},
	},
}

var replacementsInlineTargetObjectAttributes = map[string]schema.Attribute{
	"group": schema.StringAttribute{
		Optional:    true,
		Description: "The group of the referent",
	},
	"version": schema.StringAttribute{
		Optional:    true,
		Description: "The version of the referent",
	},
	"kind": schema.StringAttribute{
		Optional:    true,
		Description: "The kind of the referent",
	},
	"name": schema.StringAttribute{
		Optional:    true,
		Description: "The name of the referent",
	},
	"namespace": schema.StringAttribute{
		Optional:    true,
		Description: "The namespace of the referent",
	},
}