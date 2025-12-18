// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkDhcpIpAddressDetailParam GetL3NetworkDhcpIpAddress detail param
type GetL3NetworkDhcpIpAddressDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
}

// GetL3NetworkDhcpIpAddressParam GetL3NetworkDhcpIpAddress request param
type GetL3NetworkDhcpIpAddressParam struct {
	BaseParam
	Params GetL3NetworkDhcpIpAddressDetailParam `json:"params"`
}
