// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecurityMachineStateDetailParam ChangeSecurityMachineState detail param
type ChangeSecurityMachineStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSecurityMachineStateParam ChangeSecurityMachineState request param
type ChangeSecurityMachineStateParam struct {
	BaseParam
	Params ChangeSecurityMachineStateDetailParam `json:"params"`
}
