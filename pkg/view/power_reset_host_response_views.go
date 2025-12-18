// Copyright (c) ZStack.io, Inc.

package view

// PowerResetHostEventView PowerResetHostEvent
type PowerResetHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

