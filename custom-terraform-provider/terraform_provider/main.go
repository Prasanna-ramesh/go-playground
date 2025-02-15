package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"log"
	"terraform-provider-mockupstream/provider"
)

func main() {

	opts := providerserver.ServeOpts{
		// NOTE: This is not a typical Terraform Registry provider address,
		// such as registry.terraform.io/hashicorp/hashicups. This specific
		// provider address is used in conjunction with a specific Terraform
		// CLI configuration for manual development testing of this provider.
		// Refer to README for configuring this locally
		// This follows the pattern [<HOSTNAME>/]<NAMESPACE>/<TYPE>
		Address: "terraform.local/prasanna-ramesh/mockupstream",
		Debug:   false,
	}

	err := providerserver.Serve(context.Background(), provider.New("dev"), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
