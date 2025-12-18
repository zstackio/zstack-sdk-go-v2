// Copyright (c) ZStack.io, Inc.

package view

// GetHostNetworkFactsView GetHostNetworkFacts
type GetHostNetworkFactsView struct {
	Bondings []HostNetworkBondingInventoryView `json:"bondings,omitempty"`
	Nics []HostNetworkInterfaceInventoryView `json:"nics,omitempty"`
	Success bool `json:"success,omitempty"`
}

