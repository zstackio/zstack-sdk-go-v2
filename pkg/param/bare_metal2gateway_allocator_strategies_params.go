// Copyright (c) ZStack.io, Inc.

package param

// GetBareMetal2GatewayAllocatorStrategiesDetailParam GetBareMetal2GatewayAllocatorStrategies详细参数
type GetBareMetal2GatewayAllocatorStrategiesDetailParam struct {
}

// GetBareMetal2GatewayAllocatorStrategiesParam GetBareMetal2GatewayAllocatorStrategies请求参数
type GetBareMetal2GatewayAllocatorStrategiesParam struct {
	BaseParam
	Params GetBareMetal2GatewayAllocatorStrategiesDetailParam `json:"params"` // 详细参数
}

