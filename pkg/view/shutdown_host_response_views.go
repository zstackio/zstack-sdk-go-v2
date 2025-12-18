// Copyright (c) ZStack.io, Inc.

package view

// ShutdownHostEventView ShutdownHostEvent
type ShutdownHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

