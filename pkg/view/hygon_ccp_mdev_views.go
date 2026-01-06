// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HygonCcpMdevInventoryView HygonCcpMdev
type HygonCcpMdevInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MdevUuid string `json:"mdevUuid,omitempty"`
	CcpDeviceUuid string `json:"ccpDeviceUuid,omitempty"`
	VendorIdx int `json:"vendorIdx,omitempty"`
	UseFlag int `json:"useFlag,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

