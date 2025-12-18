// Copyright (c) ZStack.io, Inc.

package param

// GetGlobalConfigOptionsDetailParam GetGlobalConfigOptions详细参数
type GetGlobalConfigOptionsDetailParam struct {
	rest string `json:"category" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
}

// GetGlobalConfigOptionsParam GetGlobalConfigOptions请求参数
type GetGlobalConfigOptionsParam struct {
	BaseParam
	Params GetGlobalConfigOptionsDetailParam `json:"params"` // 详细参数
}

