// Copyright (c) ZStack.io, Inc.

package param

// AddReservedIpRangeDetailParam AddReservedIpRange详细参数
type AddReservedIpRangeDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"startIp" validate:"required"` // 必填
	rest string `json:"endIp" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddReservedIpRangeParam AddReservedIpRange请求参数
type AddReservedIpRangeParam struct {
	BaseParam
	Params AddReservedIpRangeDetailParam `json:"params"` // 详细参数
}

