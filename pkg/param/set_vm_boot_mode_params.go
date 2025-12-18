// Copyright (c) ZStack.io, Inc.

package param

// SetVmBootModeDetailParam SetVmBootMode详细参数
type SetVmBootModeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"bootMode" validate:"required"` // 必填
}

// SetVmBootModeParam SetVmBootMode请求参数
type SetVmBootModeParam struct {
	BaseParam
	Params SetVmBootModeDetailParam `json:"params"` // 详细参数
}

