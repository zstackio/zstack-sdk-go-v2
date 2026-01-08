// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HygonCcpDeviceInventoryView HygonCcpDevice
type HygonCcpDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid     string `json:"hostUuid,omitempty"`
	PciBdf       string `json:"pciBdf,omitempty"`
	DeviceType   string `json:"deviceType,omitempty"`
	DeviceId     string `json:"deviceId,omitempty"`
	DriverStatus string `json:"driverStatus,omitempty"`
	IsMasterPsp  bool   `json:"isMasterPsp,omitempty"`
	VendorIdx    int    `json:"vendorIdx,omitempty"`
	State        string `json:"state,omitempty"`
}

// QueryHygonDeviceView QueryHygonDevice
type QueryHygonDeviceView struct {
	Inventories []HygonCcpDeviceInventoryView `json:"inventories,omitempty"`
}
