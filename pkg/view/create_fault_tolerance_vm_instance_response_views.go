// Copyright (c) ZStack.io, Inc.

package view

// CreateFaultToleranceVmInstanceEventView CreateFaultToleranceVmInstanceEvent
type CreateFaultToleranceVmInstanceEventView struct {
	PrimaryVmInventory VmInstanceInventoryView `json:"primaryVmInventory,omitempty"`
	SecondaryVmInventory VmInstanceInventoryView `json:"secondaryVmInventory,omitempty"`
	FaultToleranceVmGroupInventory VmInstanceInventoryView `json:"faultToleranceVmGroupInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

