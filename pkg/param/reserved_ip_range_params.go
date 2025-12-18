// Copyright (c) ZStack.io, Inc.

package param

// DeleteReservedIpRangeDetailParam DeleteReservedIpRange详细参数
type DeleteReservedIpRangeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteReservedIpRangeParam DeleteReservedIpRange请求参数
type DeleteReservedIpRangeParam struct {
	BaseParam
	Params DeleteReservedIpRangeDetailParam `json:"params"` // 详细参数
}

