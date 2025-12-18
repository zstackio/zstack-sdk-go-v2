// Copyright (c) ZStack.io, Inc.

package param

// GetFreeIpDetailParam GetFreeIp detail param
type GetFreeIpDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	IpRangeUuid string `json:"ipRangeUuid,omitempty"`
	Start string `json:"start,omitempty"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	Limit int `json:"limit,omitempty"`
}

// GetFreeIpParam GetFreeIp request param
type GetFreeIpParam struct {
	BaseParam
	Params GetFreeIpDetailParam `json:"params"`
}
