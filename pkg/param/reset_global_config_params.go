// Copyright (c) ZStack.io, Inc.

package param

// ResetGlobalConfigDetailParam ResetGlobalConfig详细参数
type ResetGlobalConfigDetailParam struct {
}

// ResetGlobalConfigParam ResetGlobalConfig请求参数
type ResetGlobalConfigParam struct {
	BaseParam
	Params ResetGlobalConfigDetailParam `json:"params"` // 详细参数
}

