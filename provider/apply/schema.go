package apply

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

var kustomizeApplySchema = schema.Schema{
	Description: "This resource renders Kubernetes manifests using Kustomize and apply the generated manifests, which is equivalent to `kustomize build | kubectl apply -f`.",
	Attributes: lo.Assign[string, schema.Attribute](
		kustomizeAttributes,
		map[string]schema.Attribute{
			// Required for testing framework
			"id": schema.StringAttribute{
				Computed: true,
			},
			// Some read-only attributes
			"yaml": schema.StringAttribute{
				Computed:    true,
				Description: "The generated Kubernetes manifests in yaml format.",
			},
		},
	),
}

var kustomizeAttributes = map[string]schema.Attribute{
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
