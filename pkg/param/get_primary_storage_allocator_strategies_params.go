// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageAllocatorStrategiesDetailParam GetPrimaryStorageAllocatorStrategies detail param
type GetPrimaryStorageAllocatorStrategiesDetailParam struct {
}

// GetPrimaryStorageAllocatorStrategiesParam GetPrimaryStorageAllocatorStrategies request param
type GetPrimaryStorageAllocatorStrategiesParam struct {
	BaseParam
	Params GetPrimaryStorageAllocatorStrategiesDetailParam `json:"params"`
}
