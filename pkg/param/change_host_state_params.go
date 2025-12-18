// Copyright (c) ZStack.io, Inc.

package param

// ChangeHostStateDetailParam ChangeHostState detail param
type ChangeHostStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeHostStateParam ChangeHostState request param
type ChangeHostStateParam struct {
	BaseParam
	Params ChangeHostStateDetailParam `json:"params"`
}
