package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/thib92/terraform-provider-elastic-cloud/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.New})
}
