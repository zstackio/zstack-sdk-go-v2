// Copyright (c) ZStack.io, Inc.

package param

// GetContainerUsageDetailParam GetContainerUsage详细参数
type GetContainerUsageDetailParam struct {
}

// GetContainerUsageParam GetContainerUsage请求参数
type GetContainerUsageParam struct {
	BaseParam
	Params GetContainerUsageDetailParam `json:"params"` // 详细参数
}

