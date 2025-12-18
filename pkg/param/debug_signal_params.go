// Copyright (c) ZStack.io, Inc.

package param

// DebugSignalDetailParam DebugSignal详细参数
type DebugSignalDetailParam struct {
	rest []string `json:"signals" validate:"required"` // 必填
}

// DebugSignalParam DebugSignal请求参数
type DebugSignalParam struct {
	BaseParam
	Params DebugSignalDetailParam `json:"params"` // 详细参数
}

