// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PciDeviceOfferingInstanceOfferingRefInventoryView PciDeviceOfferingInstanceOfferingRef
type PciDeviceOfferingInstanceOfferingRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest string `json:"pciDeviceOfferingUuid,omitempty"`
	rest interface{} `json:"metadata,omitempty"`
	rest int `json:"pciDeviceCount,omitempty"`
}

