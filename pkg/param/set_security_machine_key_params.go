// Copyright (c) ZStack.io, Inc.

package param

// SetSecurityMachineKeyDetailParam SetSecurityMachineKey detail param
type SetSecurityMachineKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	TokenName string `json:"tokenName" validate:"required"`
	DryRun bool `json:"dryRun,omitempty"`
}

// SetSecurityMachineKeyParam SetSecurityMachineKey request param
type SetSecurityMachineKeyParam struct {
	BaseParam
	Params SetSecurityMachineKeyDetailParam `json:"params"`
}
