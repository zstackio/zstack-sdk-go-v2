// Copyright (c) ZStack.io, Inc.

package view

// DeleteVmSshKeyEventView DeleteVmSshKeyEvent
type DeleteVmSshKeyEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

