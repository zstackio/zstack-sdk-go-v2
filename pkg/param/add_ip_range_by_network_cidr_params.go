// Copyright (c) ZStack.io, Inc.

package param

// AddIpRangeByNetworkCidrDetailParam AddIpRangeByNetworkCidr detail param
type AddIpRangeByNetworkCidrDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	NetworkCidr string `json:"networkCidr" validate:"required"`
	Gateway string `json:"gateway,omitempty"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIpRangeByNetworkCidrParam AddIpRangeByNetworkCidr request param
type AddIpRangeByNetworkCidrParam struct {
	BaseParam
	Params AddIpRangeByNetworkCidrDetailParam `json:"params"`
}
