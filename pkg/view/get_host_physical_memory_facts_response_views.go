// Copyright (c) ZStack.io, Inc.

package view

// GetHostPhysicalMemoryFactsView GetHostPhysicalMemoryFacts
type GetHostPhysicalMemoryFactsView struct {
	Inventories []HostPhysicalMemoryInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

