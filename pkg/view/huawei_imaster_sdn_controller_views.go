// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterSdnControllerInventoryView HuaweiIMasterSdnController
type HuaweiIMasterSdnControllerInventoryView struct {
	Fabrics []HuaweiIMasterFabricInventoryView `json:"fabrics,omitempty"`
	Tenants []HuaweiIMasterTenantInventoryView `json:"tenants,omitempty"`
	Vpcs []HuaweiIMasterVpcInventoryView `json:"vpcs,omitempty"`
	Vrouters []HuaweiIMasterVRouterInventoryView `json:"vrouters,omitempty"`
	VlanRanges []SdnVlanRangeView `json:"vlanRanges,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VendorType *string `json:"vendorType,omitempty"`
	VendorVersion *string `json:"vendorVersion,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Status string `json:"status,omitempty"`
	HostRefs []SdnControllerHostRefInventoryView `json:"hostRefs,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

