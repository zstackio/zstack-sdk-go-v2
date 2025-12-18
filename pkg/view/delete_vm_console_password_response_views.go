// Copyright (c) ZStack.io, Inc.

package view

// DeleteVmConsolePasswordEventView DeleteVmConsolePasswordEvent
type DeleteVmConsolePasswordEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

