// Copyright (c) ZStack.io, Inc.

package param

// DeleteLogServerDetailParam DeleteLogServer detail param
type DeleteLogServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteLogServerParam DeleteLogServer request param
type DeleteLogServerParam struct {
	BaseParam
	Params DeleteLogServerDetailParam `json:"params"`
}
