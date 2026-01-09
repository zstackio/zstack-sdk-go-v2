// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HygonCcpDeviceInventoryView HygonCcpDevice
type HygonCcpDeviceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PciBdf *string `json:"pciBdf,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	DriverStatus *string `json:"driverStatus,omitempty"`
	IsMasterPsp *bool `json:"isMasterPsp,omitempty"`
	VendorIdx *int `json:"vendorIdx,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryHygonDeviceView QueryHygonDevice
type QueryHygonDeviceView struct {
	Inventories []HygonCcpDeviceInventoryView `json:"inventories,omitempty"`
}

