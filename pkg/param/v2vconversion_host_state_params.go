// Copyright (c) ZStack.io, Inc.

package param

// ChangeV2VConversionHostStateDetailParam ChangeV2VConversionHostState详细参数
type ChangeV2VConversionHostStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeV2VConversionHostStateParam ChangeV2VConversionHostState请求参数
type ChangeV2VConversionHostStateParam struct {
	BaseParam
	Params ChangeV2VConversionHostStateDetailParam `json:"params"` // 详细参数
}

