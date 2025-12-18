// Copyright (c) ZStack.io, Inc.

package view

// GetHostResourceAllocationEventView GetHostResourceAllocationEvent
type GetHostResourceAllocationEventView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VCPUPin []interface{} `json:"vCPUPin,omitempty"`
}

