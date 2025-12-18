// Copyright (c) ZStack.io, Inc.

package param

// ChangeIPSecConnectionStateDetailParam ChangeIPSecConnectionState detail param
type ChangeIPSecConnectionStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeIPSecConnectionStateParam ChangeIPSecConnectionState request param
type ChangeIPSecConnectionStateParam struct {
	BaseParam
	Params ChangeIPSecConnectionStateDetailParam `json:"params"`
}
