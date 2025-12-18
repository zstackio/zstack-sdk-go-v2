// Copyright (c) ZStack.io, Inc.

package view

// GetCandidateNetworkInterfacesView GetCandidateNetworkInterfaces
type GetCandidateNetworkInterfacesView struct {
	SlaveNames []string `json:"slaveNames,omitempty"`
	CandidateNics []HostNetworkInterfaceInventoryView `json:"candidateNics,omitempty"`
	Success bool `json:"success,omitempty"`
}

