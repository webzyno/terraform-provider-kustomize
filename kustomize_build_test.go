package main

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccKustomizeBuild_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6Providers,
		Steps: []resource.TestStep{
			{
				Config: `data "kustomize_build" "test" {
						resources = ["github.com/hetznercloud/csi-driver/deploy/kubernetes"]
					}`,
			},
		},
	})
}
