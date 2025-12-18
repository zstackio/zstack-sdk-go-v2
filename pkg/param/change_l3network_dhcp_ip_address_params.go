// Copyright (c) ZStack.io, Inc.

package param

// ChangeL3NetworkDhcpIpAddressDetailParam ChangeL3NetworkDhcpIpAddress detail param
type ChangeL3NetworkDhcpIpAddressDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	DhcpServerIp string `json:"dhcpServerIp,omitempty"`
	Dhcpv6ServerIp string `json:"dhcpv6ServerIp,omitempty"`
}

// ChangeL3NetworkDhcpIpAddressParam ChangeL3NetworkDhcpIpAddress request param
type ChangeL3NetworkDhcpIpAddressParam struct {
	BaseParam
	Params ChangeL3NetworkDhcpIpAddressDetailParam `json:"params"`
}
