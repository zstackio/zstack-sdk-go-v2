// Copyright (c) ZStack.io, Inc.

package view

// PowerOnHostEventView PowerOnHostEvent
type PowerOnHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

