// Copyright (c) ZStack.io, Inc.

package param

// AddIpv6RangeDetailParam AddIpv6Range detail param
type AddIpv6RangeDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	StartIp string `json:"startIp" validate:"required"`
	EndIp string `json:"endIp" validate:"required"`
	Gateway string `json:"gateway" validate:"required"`
	PrefixLen int `json:"prefixLen" validate:"required"`
	AddressMode string `json:"addressMode" validate:"required"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpv6RangeParam AddIpv6Range request param
type AddIpv6RangeParam struct {
	BaseParam
	Params AddIpv6RangeDetailParam `json:"params"`
}
