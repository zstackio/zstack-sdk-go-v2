// Copyright (c) ZStack.io, Inc.

package param

// SetVmConsoleModeDetailParam SetVmConsoleMode详细参数
type SetVmConsoleModeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"mode" validate:"required"` // 必填
}

// SetVmConsoleModeParam SetVmConsoleMode请求参数
type SetVmConsoleModeParam struct {
	BaseParam
	Params SetVmConsoleModeDetailParam `json:"params"` // 详细参数
}

