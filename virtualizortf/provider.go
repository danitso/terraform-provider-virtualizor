/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package virtualizortf

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	dvProviderEndpoint = ""
	dvProviderKey      = ""

	mkProviderEndpoint = "endpoint"
	mkProviderKey      = "key"
)

// Provider returns the object for this provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc:  providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
		Schema: map[string]*schema.Schema{
			mkProviderEndpoint: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API endpoint",
			},
			mkProviderKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API key",
			},
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return nil, errors.New("Not implemented")
}
