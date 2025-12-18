// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyRouteTableInventoryView PolicyRouteTable
type PolicyRouteTableInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int `json:"tableNumber,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []PolicyRouteTableRouteEntryInventoryView `json:"routes,omitempty"`
}

