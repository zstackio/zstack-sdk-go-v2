// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PciDeviceOfferingInventoryView PciDeviceOffering
type PciDeviceOfferingInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	RamSize string `json:"ramSize,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AttachedInstanceOfferings []PciDeviceOfferingInstanceOfferingRefInventoryView `json:"attachedInstanceOfferings,omitempty"`
	MatchedPciDevices []PciDevicePciDeviceOfferingRefInventoryView `json:"matchedPciDevices,omitempty"`
}

