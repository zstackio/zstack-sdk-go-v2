// Copyright (c) ZStack.io, Inc.

package param

// GetLoginProceduresDetailParam GetLoginProcedures详细参数
type GetLoginProceduresDetailParam struct {
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"loginType" validate:"required"` // 必填
}

// GetLoginProceduresParam GetLoginProcedures请求参数
type GetLoginProceduresParam struct {
	BaseParam
	Params GetLoginProceduresDetailParam `json:"params"` // 详细参数
}

