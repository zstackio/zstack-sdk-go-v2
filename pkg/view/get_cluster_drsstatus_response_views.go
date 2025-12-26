// Copyright (c) ZStack.io, Inc.

package view

// GetClusterDRSStatusView GetClusterDRSStatus
type GetClusterDRSStatusView struct {
	HostLoadOverThreshold []HostLoadView `json:"hostLoadOverThreshold,omitempty"`
	Success bool `json:"success,omitempty"`
}

