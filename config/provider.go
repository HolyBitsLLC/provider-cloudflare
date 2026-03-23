package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	accessCluster "github.com/holybitsllc/provider-cloudflare/config/cluster/access"
	dnsCluster "github.com/holybitsllc/provider-cloudflare/config/cluster/dns"
	tunnelCluster "github.com/holybitsllc/provider-cloudflare/config/cluster/tunnel"

	accessNamespaced "github.com/holybitsllc/provider-cloudflare/config/namespaced/access"
	dnsNamespaced "github.com/holybitsllc/provider-cloudflare/config/namespaced/dns"
	tunnelNamespaced "github.com/holybitsllc/provider-cloudflare/config/namespaced/tunnel"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/holybitsllc/provider-cloudflare"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudflare."),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		accessCluster.Configure,
		tunnelCluster.Configure,
		dnsCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudflare.m."),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		accessNamespaced.Configure,
		tunnelNamespaced.Configure,
		dnsNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
