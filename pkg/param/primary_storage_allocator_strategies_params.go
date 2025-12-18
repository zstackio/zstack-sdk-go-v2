// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageAllocatorStrategiesDetailParam GetPrimaryStorageAllocatorStrategies详细参数
type GetPrimaryStorageAllocatorStrategiesDetailParam struct {
}

// GetPrimaryStorageAllocatorStrategiesParam GetPrimaryStorageAllocatorStrategies请求参数
type GetPrimaryStorageAllocatorStrategiesParam struct {
	BaseParam
	Params GetPrimaryStorageAllocatorStrategiesDetailParam `json:"params"` // 详细参数
}

