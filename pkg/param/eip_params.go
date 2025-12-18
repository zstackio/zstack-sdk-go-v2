// Copyright (c) ZStack.io, Inc.

package param

// UpdateEipDetailParam UpdateEip详细参数
type UpdateEipDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateEipParam UpdateEip请求参数
type UpdateEipParam struct {
	BaseParam
	Params UpdateEipDetailParam `json:"params"` // 详细参数
}

