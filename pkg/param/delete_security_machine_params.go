// Copyright (c) ZStack.io, Inc.

package param

// DeleteSecurityMachineDetailParam DeleteSecurityMachine detail param
type DeleteSecurityMachineDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSecurityMachineParam DeleteSecurityMachine request param
type DeleteSecurityMachineParam struct {
	BaseParam
	Params DeleteSecurityMachineDetailParam `json:"params"`
}
