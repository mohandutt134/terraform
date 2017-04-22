package main

import (
	"github.com/hashicorp/terraform/builtin/providers/papertrail"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: papertrail.Provider,
	})
}
