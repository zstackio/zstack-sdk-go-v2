// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PricePciDeviceOfferingRefInventoryView PricePciDeviceOfferingRef
type PricePciDeviceOfferingRefInventoryView struct {
	rest string `json:"priceUuid,omitempty"`
	rest string `json:"pciDeviceOfferingUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

