// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PricePciDeviceOfferingRefInventoryView PricePciDeviceOfferingRef
type PricePciDeviceOfferingRefInventoryView struct {
	PriceUuid             string    `json:"priceUuid,omitempty"`
	PciDeviceOfferingUuid string    `json:"pciDeviceOfferingUuid,omitempty"`
	CreateDate            time.Time `json:"createDate,omitempty"`
	LastOpDate            time.Time `json:"lastOpDate,omitempty"`
}
