// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BaremetalBondingInventoryView BaremetalBonding
type BaremetalBondingInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ChassisUuid string `json:"chassisUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Mode int `json:"mode,omitempty"`
	Slaves string `json:"slaves,omitempty"`
	Opts string `json:"opts,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

