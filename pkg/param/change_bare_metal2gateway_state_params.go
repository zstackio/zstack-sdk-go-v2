// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2GatewayStateDetailParam ChangeBareMetal2GatewayState detail param
type ChangeBareMetal2GatewayStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2GatewayStateParam ChangeBareMetal2GatewayState request param
type ChangeBareMetal2GatewayStateParam struct {
	BaseParam
	Params ChangeBareMetal2GatewayStateDetailParam `json:"params"`
}
