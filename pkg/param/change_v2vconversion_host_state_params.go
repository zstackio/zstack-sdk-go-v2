// Copyright (c) ZStack.io, Inc.

package param

// ChangeV2VConversionHostStateDetailParam ChangeV2VConversionHostState detail param
type ChangeV2VConversionHostStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeV2VConversionHostStateParam ChangeV2VConversionHostState request param
type ChangeV2VConversionHostStateParam struct {
	BaseParam
	Params ChangeV2VConversionHostStateDetailParam `json:"params"`
}
