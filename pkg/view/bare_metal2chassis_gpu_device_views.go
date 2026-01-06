// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ChassisGpuDeviceInventoryView BareMetal2ChassisGpuDevice
type BareMetal2ChassisGpuDeviceInventoryView struct {
	SerialNumber string `json:"serialNumber,omitempty"`
	Memory int64 `json:"memory,omitempty"`
	Power int64 `json:"power,omitempty"`
	IsDriverLoaded bool `json:"isDriverLoaded,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	PciDeviceAddress string `json:"pciDeviceAddress,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	IommuGroup string `json:"iommuGroup,omitempty"`
	Name string `json:"name,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Device string `json:"device,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryBareMetal2ChassisGpuDeviceView QueryBareMetal2ChassisGpuDevice
type QueryBareMetal2ChassisGpuDeviceView struct {
	Inventories []BareMetal2ChassisGpuDeviceInventoryView `json:"inventories,omitempty"`
}

