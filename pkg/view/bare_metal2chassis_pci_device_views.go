// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ChassisPciDeviceInventoryView BareMetal2ChassisPciDevice
type BareMetal2ChassisPciDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
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

// QueryBareMetal2ChassisPciDeviceView QueryBareMetal2ChassisPciDevice
type QueryBareMetal2ChassisPciDeviceView struct {
	Inventories []BareMetal2ChassisPciDeviceInventoryView `json:"inventories,omitempty"`
}

// UpdateBareMetal2ChassisPciDeviceEventView UpdateBareMetal2ChassisPciDeviceEvent
type UpdateBareMetal2ChassisPciDeviceEventView struct {
	Inventory BareMetal2ChassisPciDeviceInventoryView `json:"inventory,omitempty"`
}
