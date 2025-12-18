// Copyright (c) ZStack.io, Inc.

package param

// GetIpAddressCapacityDetailParam GetIpAddressCapacity详细参数
type GetIpAddressCapacityDetailParam struct {
	rest []string `json:"zoneUuids,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest []string `json:"ipRangeUuids,omitempty"`
	rest bool `json:"all,omitempty"`
}

// GetIpAddressCapacityParam GetIpAddressCapacity请求参数
type GetIpAddressCapacityParam struct {
	BaseParam
	Params GetIpAddressCapacityDetailParam `json:"params"` // 详细参数
}

