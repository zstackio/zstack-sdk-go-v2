// Copyright (c) ZStack.io, Inc.

package param

// ShutdownHostDetailParam ShutdownHost detail param
type ShutdownHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ReturnEarly bool `json:"returnEarly,omitempty"`
	Force bool `json:"force,omitempty"`
	Method string `json:"method,omitempty"`
}

// ShutdownHostParam ShutdownHost request param
type ShutdownHostParam struct {
	BaseParam
	Params ShutdownHostDetailParam `json:"params"`
}
