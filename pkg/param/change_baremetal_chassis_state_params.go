// Copyright (c) ZStack.io, Inc.

package param

// ChangeBaremetalChassisStateDetailParam ChangeBaremetalChassisState detail param
type ChangeBaremetalChassisStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBaremetalChassisStateParam ChangeBaremetalChassisState request param
type ChangeBaremetalChassisStateParam struct {
	BaseParam
	Params ChangeBaremetalChassisStateDetailParam `json:"params"`
}
