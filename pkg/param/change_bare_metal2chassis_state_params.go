// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2ChassisStateDetailParam ChangeBareMetal2ChassisState detail param
type ChangeBareMetal2ChassisStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2ChassisStateParam ChangeBareMetal2ChassisState request param
type ChangeBareMetal2ChassisStateParam struct {
	BaseParam
	Params ChangeBareMetal2ChassisStateDetailParam `json:"params"`
}
