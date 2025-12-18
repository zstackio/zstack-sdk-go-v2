// Copyright (c) ZStack.io, Inc.

package view

// GetEipAttachableVmNicsView GetEipAttachableVmNics
type GetEipAttachableVmNicsView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
	Start int `json:"start,omitempty"`
	More bool `json:"more,omitempty"`
	Success bool `json:"success,omitempty"`
}

