package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/webzyno/terraform-provider-kustomize/provider"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/webzyno/kustomize",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), provider.NewProvider, opts); err != nil {
		log.Fatal(err.Error())
	}
}
