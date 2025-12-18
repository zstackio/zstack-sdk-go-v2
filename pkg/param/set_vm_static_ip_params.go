// Copyright (c) ZStack.io, Inc.

package param

// SetVmStaticIpDetailParam SetVmStaticIp详细参数
type SetVmStaticIpDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"ip,omitempty"`
	rest string `json:"ip6,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"ipv6Gateway,omitempty"`
	rest string `json:"ipv6Prefix,omitempty"`
}

// SetVmStaticIpParam SetVmStaticIp请求参数
type SetVmStaticIpParam struct {
	BaseParam
	Params SetVmStaticIpDetailParam `json:"params"` // 详细参数
}

