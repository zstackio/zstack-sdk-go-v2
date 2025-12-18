// Copyright (c) ZStack.io, Inc.

package param

// SetVmEmulatorPinningDetailParam SetVmEmulatorPinning详细参数
type SetVmEmulatorPinningDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"emulatorPinning" validate:"required"` // 必填
}

// SetVmEmulatorPinningParam SetVmEmulatorPinning请求参数
type SetVmEmulatorPinningParam struct {
	BaseParam
	Params SetVmEmulatorPinningDetailParam `json:"params"` // 详细参数
}

