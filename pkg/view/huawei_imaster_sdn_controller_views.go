// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HuaweiIMasterSdnControllerInventoryView HuaweiIMasterSdnController
type HuaweiIMasterSdnControllerInventoryView struct {
	rest []HuaweiIMasterFabricInventoryView `json:"fabrics,omitempty"`
	rest []HuaweiIMasterTenantInventoryView `json:"tenants,omitempty"`
	rest []HuaweiIMasterVpcInventoryView `json:"vpcs,omitempty"`
	rest []HuaweiIMasterVRouterInventoryView `json:"vrouters,omitempty"`
	rest []interface{} `json:"vlanRanges,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"vendorType,omitempty"`
	rest string `json:"vendorVersion,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"status,omitempty"`
	rest []SdnControllerHostRefInventoryView `json:"hostRefs,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

