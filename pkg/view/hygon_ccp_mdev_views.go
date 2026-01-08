// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HygonCcpMdevInventoryView HygonCcpMdev
type HygonCcpMdevInventoryView struct {
	BaseInfoView
	BaseTimeView
	MdevUuid       string `json:"mdevUuid,omitempty"`
	CcpDeviceUuid  string `json:"ccpDeviceUuid,omitempty"`
	VendorIdx      int    `json:"vendorIdx,omitempty"`
	UseFlag        int    `json:"useFlag,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	Status         string `json:"status,omitempty"`
	State          string `json:"state,omitempty"`
	HostUuid       string `json:"hostUuid,omitempty"`
}
