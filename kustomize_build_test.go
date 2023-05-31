package main

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccKustomizeBuild_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6Providers,
		/*ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "registry.terraform.io/hashicorp/random",
			},
		},*/
		Steps: []resource.TestStep{
			{
				Config: `data "kustomize_build" "test" {
						resources = ["github.com/hetznercloud/csi-driver/deploy/kubernetes"]
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kustomize_build.test", "resources.0", "github.com/hetznercloud/csi-driver/deploy/kubernetes"),
				),
			},
		},
	})
}
