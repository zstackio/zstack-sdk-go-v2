// Copyright (c) ZStack.io, Inc.

package view

// DeleteVmStaticIpEventView DeleteVmStaticIpEvent
type DeleteVmStaticIpEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

