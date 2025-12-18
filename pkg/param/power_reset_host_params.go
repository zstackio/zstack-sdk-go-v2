// Copyright (c) ZStack.io, Inc.

package param

// PowerResetHostDetailParam PowerResetHost detail param
type PowerResetHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ReturnEarly bool `json:"returnEarly,omitempty"`
	Method string `json:"method,omitempty"`
}

// PowerResetHostParam PowerResetHost request param
type PowerResetHostParam struct {
	BaseParam
	Params PowerResetHostDetailParam `json:"params"`
}
