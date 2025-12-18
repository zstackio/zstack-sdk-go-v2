// Copyright (c) ZStack.io, Inc.

package view

// DeleteVmCdRomEventView DeleteVmCdRomEvent
type DeleteVmCdRomEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

