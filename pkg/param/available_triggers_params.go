// Copyright (c) ZStack.io, Inc.

package param

// GetAvailableTriggersDetailParam GetAvailableTriggers详细参数
type GetAvailableTriggersDetailParam struct {
}

// GetAvailableTriggersParam GetAvailableTriggers请求参数
type GetAvailableTriggersParam struct {
	BaseParam
	Params GetAvailableTriggersDetailParam `json:"params"` // 详细参数
}

