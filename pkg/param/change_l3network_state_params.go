// Copyright (c) ZStack.io, Inc.

package param

// ChangeL3NetworkStateDetailParam ChangeL3NetworkState detail param
type ChangeL3NetworkStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeL3NetworkStateParam ChangeL3NetworkState request param
type ChangeL3NetworkStateParam struct {
	BaseParam
	Params ChangeL3NetworkStateDetailParam `json:"params"`
}
