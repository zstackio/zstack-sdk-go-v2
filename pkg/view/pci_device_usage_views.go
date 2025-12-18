// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PciDeviceUsageInventoryView PciDeviceUsage
type PciDeviceUsageInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest int64 `json:"dateInLong,omitempty"`
	rest string `json:"pciDeviceUuid,omitempty"`
	rest string `json:"vendorId,omitempty"`
	rest string `json:"deviceId,omitempty"`
	rest string `json:"subvendorId,omitempty"`
	rest string `json:"subdeviceId,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"vmName,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"inventory,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

