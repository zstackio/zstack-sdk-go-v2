// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ConnectionRelationShipInventoryView ConnectionRelationShip
type ConnectionRelationShipInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"relationShips,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

