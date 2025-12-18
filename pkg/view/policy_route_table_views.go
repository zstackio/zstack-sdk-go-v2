// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteTableInventoryView PolicyRouteTable
type PolicyRouteTableInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	TableNumber int `json:"tableNumber,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Routes []PolicyRouteTableRouteEntryInventoryView `json:"routes,omitempty"`
}

