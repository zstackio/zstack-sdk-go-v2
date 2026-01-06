// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2BondingInventoryView BareMetal2Bonding
type BareMetal2BondingInventoryView struct {
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Slaves string `json:"slaves,omitempty"`
	Opts string `json:"opts,omitempty"`
	Mode int `json:"mode,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// QueryBareMetal2BondingView QueryBareMetal2Bonding
type QueryBareMetal2BondingView struct {
	Inventories []BareMetal2BondingInventoryView `json:"inventories,omitempty"`
}

// CreateBareMetal2BondingEventView CreateBareMetal2BondingEvent
type CreateBareMetal2BondingEventView struct {
	Inventory BareMetal2BondingInventoryView `json:"inventory,omitempty"`
}

