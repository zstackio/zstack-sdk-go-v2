// Copyright (c) ZStack.io, Inc.

package param

// GetVmEmulatorPinningDetailParam GetVmEmulatorPinning detail param
type GetVmEmulatorPinningDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmEmulatorPinningParam GetVmEmulatorPinning request param
type GetVmEmulatorPinningParam struct {
	BaseParam
	Params GetVmEmulatorPinningDetailParam `json:"params"`
}
