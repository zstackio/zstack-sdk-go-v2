// Copyright (c) ZStack.io, Inc.

package param

// UpdateGlobalConfigDetailParam UpdateGlobalConfig详细参数
type UpdateGlobalConfigDetailParam struct {
	rest string `json:"category" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"value,omitempty"`
}

// UpdateGlobalConfigParam UpdateGlobalConfig请求参数
type UpdateGlobalConfigParam struct {
	BaseParam
	Params UpdateGlobalConfigDetailParam `json:"params"` // 详细参数
}

