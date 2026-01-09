// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ChassisPciDeviceInventoryView BareMetal2ChassisPciDevice
type BareMetal2ChassisPciDeviceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ChassisUuid *string `json:"chassisUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	PciDeviceAddress *string `json:"pciDeviceAddress,omitempty"`
	VendorId *string `json:"vendorId,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	SubvendorId *string `json:"subvendorId,omitempty"`
	SubdeviceId *string `json:"subdeviceId,omitempty"`
	IommuGroup *string `json:"iommuGroup,omitempty"`
	Name string `json:"name,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Device *string `json:"device,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryBareMetal2ChassisPciDeviceView QueryBareMetal2ChassisPciDevice
type QueryBareMetal2ChassisPciDeviceView struct {
	Inventories []BareMetal2ChassisPciDeviceInventoryView `json:"inventories,omitempty"`
}

// UpdateBareMetal2ChassisPciDeviceEventView UpdateBareMetal2ChassisPciDeviceEvent
type UpdateBareMetal2ChassisPciDeviceEventView struct {
	Inventory BareMetal2ChassisPciDeviceInventoryView `json:"inventory,omitempty"`
}

