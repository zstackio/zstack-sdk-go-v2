// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PriceInventoryView Price
type PriceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceName,omitempty"`
	rest string `json:"resourceUnit,omitempty"`
	rest string `json:"timeUnit,omitempty"`
	rest float64 `json:"price,omitempty"`
	rest int64 `json:"dateInLong,omitempty"`
	rest int64 `json:"endDateInLong,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"tableUuid,omitempty"`
	rest []PricePciDeviceOfferingRefInventoryView `json:"pciDeviceOfferings,omitempty"`
	rest []PriceBareMetal2ChassisOfferingRefInventoryView `json:"bareMetal2VmOfferings,omitempty"`
}

