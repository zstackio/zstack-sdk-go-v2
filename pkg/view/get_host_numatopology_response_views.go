// Copyright (c) ZStack.io, Inc.

package view

// GetHostNUMATopologyEventView GetHostNUMATopologyEvent
type GetHostNUMATopologyEventView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Topology map[string]HostNUMANodeView `json:"topology,omitempty"`
}

