// Copyright (c) ZStack.io, Inc.

package param

// AddIpv6RangeByNetworkCidrDetailParam AddIpv6RangeByNetworkCidr详细参数
type AddIpv6RangeByNetworkCidrDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"networkCidr" validate:"required"` // 必填
	rest string `json:"addressMode" validate:"required"` // 必填
	rest string `json:"ipRangeType,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddIpv6RangeByNetworkCidrParam AddIpv6RangeByNetworkCidr请求参数
type AddIpv6RangeByNetworkCidrParam struct {
	BaseParam
	Params AddIpv6RangeByNetworkCidrDetailParam `json:"params"` // 详细参数
}

