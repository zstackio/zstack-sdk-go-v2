// Copyright (c) ZStack.io, Inc.

package view

// GetClusterHostNetworkFactsView GetClusterHostNetworkFacts
type GetClusterHostNetworkFactsView struct {
	Bondings []HostNetworkBondingInventoryView `json:"bondings,omitempty"`
	Nics []HostNetworkInterfaceInventoryView `json:"nics,omitempty"`
	Success bool `json:"success,omitempty"`
}

