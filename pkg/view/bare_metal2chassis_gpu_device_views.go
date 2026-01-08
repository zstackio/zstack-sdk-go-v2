// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ChassisGpuDeviceInventoryView BareMetal2ChassisGpuDevice
type BareMetal2ChassisGpuDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	SerialNumber     string `json:"serialNumber,omitempty"`
	Memory           int64  `json:"memory,omitempty"`
	Power            int64  `json:"power,omitempty"`
	IsDriverLoaded   bool   `json:"isDriverLoaded,omitempty"`
	ChassisUuid      string `json:"chassisUuid,omitempty"`
	Type             string `json:"type,omitempty"`
	PciDeviceAddress string `json:"pciDeviceAddress,omitempty"`
	VendorId         string `json:"vendorId,omitempty"`
	DeviceId         string `json:"deviceId,omitempty"`
	SubvendorId      string `json:"subvendorId,omitempty"`
	SubdeviceId      string `json:"subdeviceId,omitempty"`
	IommuGroup       string `json:"iommuGroup,omitempty"`
	Vendor           string `json:"vendor,omitempty"`
	Device           string `json:"device,omitempty"`
}

// QueryBareMetal2ChassisGpuDeviceView QueryBareMetal2ChassisGpuDevice
type QueryBareMetal2ChassisGpuDeviceView struct {
	Inventories []BareMetal2ChassisGpuDeviceInventoryView `json:"inventories,omitempty"`
}
