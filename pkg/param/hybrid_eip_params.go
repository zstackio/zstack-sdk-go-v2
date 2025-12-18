// Copyright (c) ZStack.io, Inc.

package param

// UpdateHybridEipDetailParam UpdateHybridEip详细参数
type UpdateHybridEipDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
}

// UpdateHybridEipParam UpdateHybridEip请求参数
type UpdateHybridEipParam struct {
	BaseParam
	Params UpdateHybridEipDetailParam `json:"params"` // 详细参数
}

