// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2ChassisPciDeviceInventoryView BareMetal2ChassisPciDevice
type BareMetal2ChassisPciDeviceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"pciDeviceAddress,omitempty"`
	rest string `json:"vendorId,omitempty"`
	rest string `json:"deviceId,omitempty"`
	rest string `json:"subvendorId,omitempty"`
	rest string `json:"subdeviceId,omitempty"`
	rest string `json:"iommuGroup,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"device,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

