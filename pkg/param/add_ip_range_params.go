// Copyright (c) ZStack.io, Inc.

package param

// AddIpRangeDetailParam AddIpRange详细参数
type AddIpRangeDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"startIp" validate:"required"` // 必填
	rest string `json:"endIp" validate:"required"` // 必填
	rest string `json:"netmask" validate:"required"` // 必填
	rest string `json:"gateway,omitempty"`
	rest string `json:"ipRangeType,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddIpRangeParam AddIpRange请求参数
type AddIpRangeParam struct {
	BaseParam
	Params AddIpRangeDetailParam `json:"params"` // 详细参数
}

