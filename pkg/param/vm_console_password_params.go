// Copyright (c) ZStack.io, Inc.

package param

// GetVmConsolePasswordDetailParam GetVmConsolePassword详细参数
type GetVmConsolePasswordDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmConsolePasswordParam GetVmConsolePassword请求参数
type GetVmConsolePasswordParam struct {
	BaseParam
	Params GetVmConsolePasswordDetailParam `json:"params"` // 详细参数
}

