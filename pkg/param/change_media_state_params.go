// Copyright (c) ZStack.io, Inc.

package param

// ChangeMediaStateDetailParam ChangeMediaState detail param
type ChangeMediaStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMediaStateParam ChangeMediaState request param
type ChangeMediaStateParam struct {
	BaseParam
	Params ChangeMediaStateDetailParam `json:"params"`
}
