package provider

import (
	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/resources"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:   Version,
		Name:      "github",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"organizations":   resources.Organizations(),
			"repositories":    resources.Repositories(),
			"teams":           resources.Teams(),
			"external_groups": resources.ExternalGroups(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
