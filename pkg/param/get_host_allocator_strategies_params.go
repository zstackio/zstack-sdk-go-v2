// Copyright (c) ZStack.io, Inc.

package param

// GetHostAllocatorStrategiesDetailParam GetHostAllocatorStrategies detail param
type GetHostAllocatorStrategiesDetailParam struct {
}

// GetHostAllocatorStrategiesParam GetHostAllocatorStrategies request param
type GetHostAllocatorStrategiesParam struct {
	BaseParam
	Params GetHostAllocatorStrategiesDetailParam `json:"params"`
}
