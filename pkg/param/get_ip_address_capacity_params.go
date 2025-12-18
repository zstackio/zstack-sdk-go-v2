// Copyright (c) ZStack.io, Inc.

package param

// GetIpAddressCapacityDetailParam GetIpAddressCapacity detail param
type GetIpAddressCapacityDetailParam struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	IpRangeUuids []string `json:"ipRangeUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetIpAddressCapacityParam GetIpAddressCapacity request param
type GetIpAddressCapacityParam struct {
	BaseParam
	Params GetIpAddressCapacityDetailParam `json:"params"`
}
