// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZnsControllerInventoryView ZnsController
type ZnsControllerInventoryView struct {
	BaseInfoView
	BaseTimeView
	TransportZones []ZnsTransportZoneInventoryView `json:"transportZones,omitempty"`
	Tenants []ZnsTenantInventoryView `json:"tenants,omitempty"`
	TenantRouters []ZnsTenantRouterInventoryView `json:"tenantRouters,omitempty"`
	VendorType string `json:"vendorType,omitempty"`
	VendorVersion string `json:"vendorVersion,omitempty"`
	Description string `json:"description,omitempty"`
	Ip string `json:"ip,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Status string `json:"status,omitempty"`
	HostRefs []SdnControllerHostRefInventoryView `json:"hostRefs,omitempty"`
}

