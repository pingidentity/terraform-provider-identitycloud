// Copyright © 2025 Ping Identity Corporation

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/provider"
)

// Run "go generate" to format example terraform files and generate the provider docs

// Format examples
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool
//go:generate go tool tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.NewFactory(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/pingidentity/identitycloud",
		Debug:   debug,
	})

	if err != nil {
		fmt.Println(err)
	}
}
