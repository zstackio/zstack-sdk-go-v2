// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2BondingInventoryView BareMetal2Bonding
type BareMetal2BondingInventoryView struct {
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"slaves,omitempty"`
	rest string `json:"opts,omitempty"`
	rest int `json:"mode,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
}

