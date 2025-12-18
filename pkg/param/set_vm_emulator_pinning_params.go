// Copyright (c) ZStack.io, Inc.

package param

// SetVmEmulatorPinningDetailParam SetVmEmulatorPinning detail param
type SetVmEmulatorPinningDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	EmulatorPinning string `json:"emulatorPinning" validate:"required"`
}

// SetVmEmulatorPinningParam SetVmEmulatorPinning request param
type SetVmEmulatorPinningParam struct {
	BaseParam
	Params SetVmEmulatorPinningDetailParam `json:"params"`
}
