// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PciDeviceOfferingInventoryView PciDeviceOffering
type PciDeviceOfferingInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"vendorId,omitempty"`
	rest string `json:"deviceId,omitempty"`
	rest string `json:"subvendorId,omitempty"`
	rest string `json:"subdeviceId,omitempty"`
	rest string `json:"ramSize,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []PciDeviceOfferingInstanceOfferingRefInventoryView `json:"attachedInstanceOfferings,omitempty"`
	rest []PciDevicePciDeviceOfferingRefInventoryView `json:"matchedPciDevices,omitempty"`
}

