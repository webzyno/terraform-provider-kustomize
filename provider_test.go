package main

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProtoV6Providers map[string]func() (tfprotov6.ProviderServer, error)

func init() {
	testAccProtoV6Providers = map[string]func() (tfprotov6.ProviderServer, error){
		"kustomize": providerserver.NewProtocol6WithError(NewProvider()),
	}
}
