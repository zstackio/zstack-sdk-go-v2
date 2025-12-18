// Copyright (c) ZStack.io, Inc.

package param

// ChangeL3NetworkDhcpIpAddressDetailParam ChangeL3NetworkDhcpIpAddress详细参数
type ChangeL3NetworkDhcpIpAddressDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"dhcpServerIp,omitempty"`
	rest string `json:"dhcpv6ServerIp,omitempty"`
}

// ChangeL3NetworkDhcpIpAddressParam ChangeL3NetworkDhcpIpAddress请求参数
type ChangeL3NetworkDhcpIpAddressParam struct {
	BaseParam
	Params ChangeL3NetworkDhcpIpAddressDetailParam `json:"params"` // 详细参数
}

