// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ConnectionRelationShipInventoryView ConnectionRelationShip
type ConnectionRelationShipInventoryView struct {
	BaseInfoView
	BaseTimeView
	RelationShips string `json:"relationShips,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchEventView UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchEvent
type UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchEventView struct {
	Inventory ConnectionRelationShipInventoryView `json:"inventory,omitempty"`
}

// CreateConnectionBetweenL3NetworkAndAliyunVSwitchEventView CreateConnectionBetweenL3NetworkAndAliyunVSwitchEvent
type CreateConnectionBetweenL3NetworkAndAliyunVSwitchEventView struct {
	Inventory ConnectionRelationShipInventoryView `json:"inventory,omitempty"`
}

// QueryConnectionBetweenL3NetworkAndAliyunVSwitchView QueryConnectionBetweenL3NetworkAndAliyunVSwitch
type QueryConnectionBetweenL3NetworkAndAliyunVSwitchView struct {
	Inventories []ConnectionRelationShipInventoryView `json:"inventories,omitempty"`
}

