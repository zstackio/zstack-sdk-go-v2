// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2BondingInventoryView BareMetal2Bonding
type BareMetal2BondingInventoryView struct {
	BaseInfoView
	BaseTimeView
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Slaves      string `json:"slaves,omitempty"`
	Opts        string `json:"opts,omitempty"`
	Mode        int    `json:"mode,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// QueryBareMetal2BondingView QueryBareMetal2Bonding
type QueryBareMetal2BondingView struct {
	Inventories []BareMetal2BondingInventoryView `json:"inventories,omitempty"`
}

// CreateBareMetal2BondingEventView CreateBareMetal2BondingEvent
type CreateBareMetal2BondingEventView struct {
	Inventory BareMetal2BondingInventoryView `json:"inventory,omitempty"`
}
