// Copyright (c) ZStack.io, Inc.

package param

// SetVmBootModeDetailParam SetVmBootMode detail param
type SetVmBootModeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BootMode string `json:"bootMode" validate:"required"`
}

// SetVmBootModeParam SetVmBootMode request param
type SetVmBootModeParam struct {
	BaseParam
	Params SetVmBootModeDetailParam `json:"params"`
}
