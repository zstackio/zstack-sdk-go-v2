// Copyright (c) ZStack.io, Inc.

package view

// GetVmvNUMATopologyView GetVmvNUMATopology
type GetVmvNUMATopologyView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Topology []interface{} `json:"topology,omitempty"`
}

