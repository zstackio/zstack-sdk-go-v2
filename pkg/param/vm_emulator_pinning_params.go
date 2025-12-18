// Copyright (c) ZStack.io, Inc.

package param

// GetVmEmulatorPinningDetailParam GetVmEmulatorPinning详细参数
type GetVmEmulatorPinningDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmEmulatorPinningParam GetVmEmulatorPinning请求参数
type GetVmEmulatorPinningParam struct {
	BaseParam
	Params GetVmEmulatorPinningDetailParam `json:"params"` // 详细参数
}

