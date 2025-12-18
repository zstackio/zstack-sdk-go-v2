// Copyright (c) ZStack.io, Inc.

package param

// ChangeImageStateDetailParam ChangeImageState detail param
type ChangeImageStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeImageStateParam ChangeImageState request param
type ChangeImageStateParam struct {
	BaseParam
	Params ChangeImageStateDetailParam `json:"params"`
}
