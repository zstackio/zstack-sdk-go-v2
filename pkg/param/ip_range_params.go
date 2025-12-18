// Copyright (c) ZStack.io, Inc.

package param

// DeleteIpRangeDetailParam DeleteIpRange详细参数
type DeleteIpRangeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteIpRangeParam DeleteIpRange请求参数
type DeleteIpRangeParam struct {
	BaseParam
	Params DeleteIpRangeDetailParam `json:"params"` // 详细参数
}

