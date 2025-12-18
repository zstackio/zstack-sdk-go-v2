// Copyright (c) ZStack.io, Inc.

package param

// AddIpv6RangeDetailParam AddIpv6Range详细参数
type AddIpv6RangeDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"startIp" validate:"required"` // 必填
	rest string `json:"endIp" validate:"required"` // 必填
	rest string `json:"gateway" validate:"required"` // 必填
	rest int `json:"prefixLen" validate:"required"` // 必填
	rest string `json:"addressMode" validate:"required"` // 必填
	rest string `json:"ipRangeType,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddIpv6RangeParam AddIpv6Range请求参数
type AddIpv6RangeParam struct {
	BaseParam
	Params AddIpv6RangeDetailParam `json:"params"` // 详细参数
}

