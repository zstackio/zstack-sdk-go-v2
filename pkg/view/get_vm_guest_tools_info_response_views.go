// Copyright (c) ZStack.io, Inc.

package view

// GetVmGuestToolsInfoView GetVmGuestToolsInfo
type GetVmGuestToolsInfoView struct {
	Version string `json:"version,omitempty"`
	Status string `json:"status,omitempty"`
	Features map[string]string `json:"features,omitempty"`
	Success bool `json:"success,omitempty"`
}

