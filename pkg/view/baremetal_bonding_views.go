// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalBondingInventoryView BaremetalBonding
type BaremetalBondingInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest int `json:"mode,omitempty"`
	rest string `json:"slaves,omitempty"`
	rest string `json:"opts,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

