// Copyright (c) ZStack.io, Inc.

package param

// SetVmStaticIpDetailParam SetVmStaticIp detail param
type SetVmStaticIpDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Ip string `json:"ip,omitempty"`
	Ip6 string `json:"ip6,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Ipv6Gateway string `json:"ipv6Gateway,omitempty"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
}

// SetVmStaticIpParam SetVmStaticIp request param
type SetVmStaticIpParam struct {
	BaseParam
	Params SetVmStaticIpDetailParam `json:"params"`
}
