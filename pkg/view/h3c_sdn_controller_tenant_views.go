// Copyright (c) ZStack.io, Inc.

package view

import "time"

// H3cSdnControllerTenantInventoryView H3cSdnControllerTenant
type H3cSdnControllerTenantInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"sdnControllerUuid,omitempty"`
	rest string `json:"tenantUuid,omitempty"`
	rest string `json:"vdsUuid,omitempty"`
	rest string `json:"tenantName,omitempty"`
	rest string `json:"vdsName,omitempty"`
	rest string `json:"cloudDomainName,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

