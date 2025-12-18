// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmBootModeDetailParam DeleteVmBootMode详细参数
type DeleteVmBootModeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVmBootModeParam DeleteVmBootMode请求参数
type DeleteVmBootModeParam struct {
	BaseParam
	Params DeleteVmBootModeDetailParam `json:"params"` // 详细参数
}

