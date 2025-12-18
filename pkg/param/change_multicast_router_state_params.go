// Copyright (c) ZStack.io, Inc.

package param

// ChangeMulticastRouterStateDetailParam ChangeMulticastRouterState detail param
type ChangeMulticastRouterStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeMulticastRouterStateParam ChangeMulticastRouterState request param
type ChangeMulticastRouterStateParam struct {
	BaseParam
	Params ChangeMulticastRouterStateDetailParam `json:"params"`
}
