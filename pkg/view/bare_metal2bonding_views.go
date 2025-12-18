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
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

