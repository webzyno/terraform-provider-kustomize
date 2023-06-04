package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestKustomizeApply(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6Providers,
		Steps: []resource.TestStep{
			{
				Config: `
					provider "kustomize" {
						kubernetes = {
							config_path = "~/.kube/config"
						}
					}

					resource "kustomize_apply" "metallb" {
						resources = ["github.com/metallb/metallb/config/native?ref=v0.13.10"]
						namespace = "metallb-system"
					}
					`,
				Check: resource.ComposeTestCheckFunc(),
			},
			/*{
				Config: `
					provider "kustomize" {
						kubernetes = {
							config_path = "~/.kube/config"
						}
					}

					resource "kustomize_apply" "metallb" {
						resources = ["github.com/metallb/metallb/config/native?ref=v0.13.10"]
						namespace = "metallb-system"
						common_annotations = {
							"webzyno.com/components" = "load-balancer"
						}
					}
					`,
				Check: resource.ComposeTestCheckFunc(),
			},*/
		},
	})
}
