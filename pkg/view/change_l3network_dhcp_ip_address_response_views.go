// Copyright (c) ZStack.io, Inc.

package view

// ChangeL3NetworkDhcpIpAddressEventView ChangeL3NetworkDhcpIpAddressEvent
type ChangeL3NetworkDhcpIpAddressEventView struct {
	DhcpServerIp string `json:"dhcpServerIp,omitempty"`
	Dhcpv6ServerIp string `json:"dhcpv6ServerIp,omitempty"`
	Success bool `json:"success,omitempty"`
}

