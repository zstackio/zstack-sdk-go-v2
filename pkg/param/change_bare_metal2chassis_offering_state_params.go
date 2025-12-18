// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2ChassisOfferingStateDetailParam ChangeBareMetal2ChassisOfferingState detail param
type ChangeBareMetal2ChassisOfferingStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeBareMetal2ChassisOfferingStateParam ChangeBareMetal2ChassisOfferingState request param
type ChangeBareMetal2ChassisOfferingStateParam struct {
	BaseParam
	Params ChangeBareMetal2ChassisOfferingStateDetailParam `json:"params"`
}
