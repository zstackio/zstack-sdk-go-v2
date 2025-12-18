// Copyright (c) ZStack.io, Inc.

package param

// AddIpRangeDetailParam AddIpRange detail param
type AddIpRangeDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	StartIp string `json:"startIp" validate:"required"`
	EndIp string `json:"endIp" validate:"required"`
	Netmask string `json:"netmask" validate:"required"`
	Gateway string `json:"gateway,omitempty"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpRangeParam AddIpRange request param
type AddIpRangeParam struct {
	BaseParam
	Params AddIpRangeDetailParam `json:"params"`
}
