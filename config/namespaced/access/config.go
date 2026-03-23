package access

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_access_organization", func(r *config.Resource) {
		r.ShortGroup = "access"
	})
	p.AddResourceConfigurator("cloudflare_access_identity_provider", func(r *config.Resource) {
		r.ShortGroup = "access"
		// The upstream TF provider uses StateFunc instead of Sensitive: true for these
		// fields, so upjet doesn't generate SecretKeyRef variants by default. We override
		// to ensure credentials are never stored in plaintext in the CRD/etcd.
		if s, ok := r.TerraformResource.Schema["config"]; ok {
			if elem, ok := s.Elem.(*schema.Resource); ok {
				if f, ok := elem.Schema["client_secret"]; ok {
					f.Sensitive = true
				}
				if f, ok := elem.Schema["api_token"]; ok {
					f.Sensitive = true
				}
				if f, ok := elem.Schema["idp_public_cert"]; ok {
					f.Sensitive = true
				}
			}
		}
	})
	p.AddResourceConfigurator("cloudflare_access_group", func(r *config.Resource) {
		r.ShortGroup = "access"
	})
	p.AddResourceConfigurator("cloudflare_access_application", func(r *config.Resource) {
		r.ShortGroup = "access"
	})
	p.AddResourceConfigurator("cloudflare_access_policy", func(r *config.Resource) {
		r.ShortGroup = "access"
	})
}
