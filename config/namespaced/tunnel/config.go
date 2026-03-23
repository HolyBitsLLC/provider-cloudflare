package tunnel

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_tunnel", func(r *config.Resource) {
		r.ShortGroup = "tunnel"
	})
	p.AddResourceConfigurator("cloudflare_tunnel_config", func(r *config.Resource) {
		r.ShortGroup = "tunnel"
		r.Kind = "TunnelConfig"
	})
	p.AddResourceConfigurator("cloudflare_tunnel_route", func(r *config.Resource) {
		r.ShortGroup = "tunnel"
	})
}
