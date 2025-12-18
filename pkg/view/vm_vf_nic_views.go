// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmVfNicInventoryView VmVfNic
type VmVfNicInventoryView struct {
	rest string `json:"pciDeviceUuid,omitempty"`
	rest string `json:"haState,omitempty"`
	rest string `json:"secondaryPciDeviceUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"mac,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"metaData,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"driverType,omitempty"`
	rest []UsedIpInventoryView `json:"usedIps,omitempty"`
	rest string `json:"internalName,omitempty"`
	rest int `json:"deviceId,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

