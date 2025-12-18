// Copyright (c) ZStack.io, Inc.

package param

// PowerOnHostDetailParam PowerOnHost detail param
type PowerOnHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ReturnEarly bool `json:"returnEarly,omitempty"`
}

// PowerOnHostParam PowerOnHost request param
type PowerOnHostParam struct {
	BaseParam
	Params PowerOnHostDetailParam `json:"params"`
}
