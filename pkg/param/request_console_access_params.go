// Copyright (c) ZStack.io, Inc.

package param

// RequestConsoleAccessDetailParam RequestConsoleAccess详细参数
type RequestConsoleAccessDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// RequestConsoleAccessParam RequestConsoleAccess请求参数
type RequestConsoleAccessParam struct {
	BaseParam
	Params RequestConsoleAccessDetailParam `json:"params"` // 详细参数
}

