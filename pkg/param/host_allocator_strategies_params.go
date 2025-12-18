// Copyright (c) ZStack.io, Inc.

package param

// GetHostAllocatorStrategiesDetailParam GetHostAllocatorStrategies详细参数
type GetHostAllocatorStrategiesDetailParam struct {
}

// GetHostAllocatorStrategiesParam GetHostAllocatorStrategies请求参数
type GetHostAllocatorStrategiesParam struct {
	BaseParam
	Params GetHostAllocatorStrategiesDetailParam `json:"params"` // 详细参数
}

