// Copyright (c) ZStack.io, Inc.

package param

// DeleteVniRangeDetailParam DeleteVniRange详细参数
type DeleteVniRangeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVniRangeParam DeleteVniRange请求参数
type DeleteVniRangeParam struct {
	BaseParam
	Params DeleteVniRangeDetailParam `json:"params"` // 详细参数
}

