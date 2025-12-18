// Copyright (c) ZStack.io, Inc.

package param

// DeleteV2VConversionHostDetailParam DeleteV2VConversionHost detail param
type DeleteV2VConversionHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteV2VConversionHostParam DeleteV2VConversionHost request param
type DeleteV2VConversionHostParam struct {
	BaseParam
	Params DeleteV2VConversionHostDetailParam `json:"params"`
}
