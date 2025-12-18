// Copyright (c) ZStack.io, Inc.

package view

// UpdateSecurityMachineEventView UpdateSecurityMachineEvent
type UpdateSecurityMachineEventView struct {
	Inventory SecurityMachineInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

