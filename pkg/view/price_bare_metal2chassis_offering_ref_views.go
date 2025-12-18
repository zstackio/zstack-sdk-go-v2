// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PriceBareMetal2ChassisOfferingRefInventoryView PriceBareMetal2ChassisOfferingRef
type PriceBareMetal2ChassisOfferingRefInventoryView struct {
	rest string `json:"priceUuid,omitempty"`
	rest string `json:"bareMetal2ChassisOfferingUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

