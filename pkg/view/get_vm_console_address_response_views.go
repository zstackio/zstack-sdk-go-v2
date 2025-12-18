// Copyright (c) ZStack.io, Inc.

package view

// GetVmConsoleAddressView GetVmConsoleAddress
type GetVmConsoleAddressView struct {
	HostIp string `json:"hostIp,omitempty"`
	Port int `json:"port,omitempty"`
	Path string `json:"path,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	VdiPortInfo interface{} `json:"vdiPortInfo,omitempty"`
	Success bool `json:"success,omitempty"`
}

