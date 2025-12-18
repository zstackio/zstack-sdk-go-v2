// Copyright (c) ZStack.io, Inc.

package param

// DeleteV2VConversionHostDetailParam DeleteV2VConversionHost详细参数
type DeleteV2VConversionHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteV2VConversionHostParam DeleteV2VConversionHost请求参数
type DeleteV2VConversionHostParam struct {
	BaseParam
	Params DeleteV2VConversionHostDetailParam `json:"params"` // 详细参数
}

