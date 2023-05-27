package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/webzyno/kustomize",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), nil, opts); err != nil {
		log.Fatal(err.Error())
	}
}
