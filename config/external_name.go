package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Access
	"cloudflare_access_organization":      config.IdentifierFromProvider,
	"cloudflare_access_identity_provider": config.IdentifierFromProvider,
	"cloudflare_access_group":             config.IdentifierFromProvider,
	"cloudflare_access_application":       config.IdentifierFromProvider,
	"cloudflare_access_policy":            config.IdentifierFromProvider,

	// Tunnel
	"cloudflare_tunnel_route":   config.IdentifierFromProvider,

	// DNS
	"cloudflare_record": config.IdentifierFromProvider,
	"cloudflare_zone":   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
