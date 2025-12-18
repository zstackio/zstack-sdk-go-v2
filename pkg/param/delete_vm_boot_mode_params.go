// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmBootModeDetailParam DeleteVmBootMode detail param
type DeleteVmBootModeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmBootModeParam DeleteVmBootMode request param
type DeleteVmBootModeParam struct {
	BaseParam
	Params DeleteVmBootModeDetailParam `json:"params"`
}
