// Copyright (c) ZStack.io, Inc.

package view

// ExecuteGuestVmCommandEventView ExecuteGuestVmCommandEvent
type ExecuteGuestVmCommandEventView struct {
	Stream string `json:"stream,omitempty"`
	VmInstance VmInstanceInventoryView `json:"vmInstance,omitempty"`
}

