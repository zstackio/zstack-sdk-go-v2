// Copyright (c) ZStack.io, Inc.

package param

// SetVmConsolePasswordDetailParam SetVmConsolePassword详细参数
type SetVmConsolePasswordDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"consolePassword" validate:"required"` // 必填
}

// SetVmConsolePasswordParam SetVmConsolePassword请求参数
type SetVmConsolePasswordParam struct {
	BaseParam
	Params SetVmConsolePasswordDetailParam `json:"params"` // 详细参数
}

