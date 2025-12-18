// Copyright (c) ZStack.io, Inc.

package param

// GetCurrentTimeDetailParam GetCurrentTime详细参数
type GetCurrentTimeDetailParam struct {
}

// GetCurrentTimeParam GetCurrentTime请求参数
type GetCurrentTimeParam struct {
	BaseParam
	Params GetCurrentTimeDetailParam `json:"params"` // 详细参数
}

