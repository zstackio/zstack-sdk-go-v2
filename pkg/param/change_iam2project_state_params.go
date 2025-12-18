// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2ProjectStateDetailParam ChangeIAM2ProjectState detail param
type ChangeIAM2ProjectStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2ProjectStateParam ChangeIAM2ProjectState request param
type ChangeIAM2ProjectStateParam struct {
	BaseParam
	Params ChangeIAM2ProjectStateDetailParam `json:"params"`
}
