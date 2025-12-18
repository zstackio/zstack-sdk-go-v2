// Copyright (c) ZStack.io, Inc.

package param

// UpdateConsolePasswordDetailParam UpdateConsolePassword详细参数
type UpdateConsolePasswordDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
}

// UpdateConsolePasswordParam UpdateConsolePassword请求参数
type UpdateConsolePasswordParam struct {
	BaseParam
	Params UpdateConsolePasswordDetailParam `json:"params"` // 详细参数
}

