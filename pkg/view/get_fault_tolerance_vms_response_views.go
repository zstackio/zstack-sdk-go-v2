// Copyright (c) ZStack.io, Inc.

package view

// GetFaultToleranceVmsView GetFaultToleranceVms
type GetFaultToleranceVmsView struct {
	PrimaryVmInventory VmInstanceInventoryView `json:"primaryVmInventory,omitempty"`
	SecondaryVmInventory VmInstanceInventoryView `json:"secondaryVmInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

