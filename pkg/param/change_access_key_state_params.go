// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccessKeyStateDetailParam ChangeAccessKeyState detail param
type ChangeAccessKeyStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAccessKeyStateParam ChangeAccessKeyState request param
type ChangeAccessKeyStateParam struct {
	BaseParam
	Params ChangeAccessKeyStateDetailParam `json:"params"`
}
