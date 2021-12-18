package resources

import (
	"embed"
	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-github/client"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*.sql
	providerMigrations embed.FS
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "github",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"repositories": Repositories(),
		},
		Migrations: providerMigrations,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
