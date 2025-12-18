// Copyright (c) ZStack.io, Inc.

package param

// GetFreeIpDetailParam GetFreeIp详细参数
type GetFreeIpDetailParam struct {
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"ipRangeUuid,omitempty"`
	rest string `json:"start,omitempty"`
	rest string `json:"ipRangeType,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest int `json:"limit,omitempty"`
}

// GetFreeIpParam GetFreeIp请求参数
type GetFreeIpParam struct {
	BaseParam
	Params GetFreeIpDetailParam `json:"params"` // 详细参数
}

