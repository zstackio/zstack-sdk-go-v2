// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PciDeviceMdevSpecRefInventoryView PciDeviceMdevSpecRef
type PciDeviceMdevSpecRefInventoryView struct {
	rest string `json:"pciDeviceUuid,omitempty"`
	rest string `json:"mdevSpecUuid,omitempty"`
	rest bool `json:"effective,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

