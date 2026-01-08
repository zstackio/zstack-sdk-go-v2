// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PriceBareMetal2ChassisOfferingRefInventoryView PriceBareMetal2ChassisOfferingRef
type PriceBareMetal2ChassisOfferingRefInventoryView struct {
	PriceUuid                     string    `json:"priceUuid,omitempty"`
	BareMetal2ChassisOfferingUuid string    `json:"bareMetal2ChassisOfferingUuid,omitempty"`
	CreateDate                    time.Time `json:"createDate,omitempty"`
	LastOpDate                    time.Time `json:"lastOpDate,omitempty"`
}
