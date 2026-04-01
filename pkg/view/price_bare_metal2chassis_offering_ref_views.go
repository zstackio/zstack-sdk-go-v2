// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PriceBareMetal2ChassisOfferingRefInventoryView PriceBareMetal2ChassisOfferingRef
type PriceBareMetal2ChassisOfferingRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	PriceUuid string `json:"priceUuid,omitempty"`
	BareMetal2ChassisOfferingUuid string `json:"bareMetal2ChassisOfferingUuid,omitempty"`
}

