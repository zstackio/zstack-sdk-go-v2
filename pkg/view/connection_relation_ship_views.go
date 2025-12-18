// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ConnectionRelationShipInventoryView ConnectionRelationShip
type ConnectionRelationShipInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	RelationShips string `json:"relationShips,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

