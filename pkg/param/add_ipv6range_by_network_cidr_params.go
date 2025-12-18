// Copyright (c) ZStack.io, Inc.

package param

// AddIpv6RangeByNetworkCidrDetailParam AddIpv6RangeByNetworkCidr detail param
type AddIpv6RangeByNetworkCidrDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	NetworkCidr string `json:"networkCidr" validate:"required"`
	AddressMode string `json:"addressMode" validate:"required"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpv6RangeByNetworkCidrParam AddIpv6RangeByNetworkCidr request param
type AddIpv6RangeByNetworkCidrParam struct {
	BaseParam
	Params AddIpv6RangeByNetworkCidrDetailParam `json:"params"`
}
