// Copyright (c) ZStack.io, Inc.

package view

// FailoverFaultToleranceVmEventView FailoverFaultToleranceVmEvent
type FailoverFaultToleranceVmEventView struct {
	PrimaryVmInventory VmInstanceInventoryView `json:"primaryVmInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

