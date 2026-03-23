// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/access/application"
	group "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/access/group"
	identityprovider "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/access/identityprovider"
	organization "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/access/organization"
	policy "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/access/policy"
	record "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/dns/record"
	zone "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/dns/zone"
	providerconfig "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/providerconfig"
	route "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/tunnel/route"
	tunnel "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/tunnel/tunnel"
	tunnelconfig "github.com/holybitsllc/provider-cloudflare/internal/controller/cluster/tunnel/tunnelconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		group.Setup,
		identityprovider.Setup,
		organization.Setup,
		policy.Setup,
		record.Setup,
		zone.Setup,
		providerconfig.Setup,
		route.Setup,
		tunnel.Setup,
		tunnelconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		group.SetupGated,
		identityprovider.SetupGated,
		organization.SetupGated,
		policy.SetupGated,
		record.SetupGated,
		zone.SetupGated,
		providerconfig.SetupGated,
		route.SetupGated,
		tunnel.SetupGated,
		tunnelconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
