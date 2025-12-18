// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2VirtualIDStateDetailParam ChangeIAM2VirtualIDState detail param
type ChangeIAM2VirtualIDStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIAM2VirtualIDStateParam ChangeIAM2VirtualIDState request param
type ChangeIAM2VirtualIDStateParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDStateDetailParam `json:"params"`
}
