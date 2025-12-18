// Copyright (c) ZStack.io, Inc.

package param

// ChangeEipStateDetailParam ChangeEipState detail param
type ChangeEipStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeEipStateParam ChangeEipState request param
type ChangeEipStateParam struct {
	BaseParam
	Params ChangeEipStateDetailParam `json:"params"`
}
