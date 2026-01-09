// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// H3cSdnControllerTenantInventoryView H3cSdnControllerTenant
type H3cSdnControllerTenantInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	SdnControllerUuid *string `json:"sdnControllerUuid,omitempty"`
	TenantUuid *string `json:"tenantUuid,omitempty"`
	VdsUuid *string `json:"vdsUuid,omitempty"`
	TenantName *string `json:"tenantName,omitempty"`
	VdsName *string `json:"vdsName,omitempty"`
	CloudDomainName *string `json:"cloudDomainName,omitempty"`
	State *string `json:"state,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

