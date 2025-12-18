// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2VirtualIDGroupStateDetailParam ChangeIAM2VirtualIDGroupState detail param
type ChangeIAM2VirtualIDGroupStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2VirtualIDGroupStateParam ChangeIAM2VirtualIDGroupState request param
type ChangeIAM2VirtualIDGroupStateParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDGroupStateDetailParam `json:"params"`
}
