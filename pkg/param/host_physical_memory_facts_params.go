// Copyright (c) ZStack.io, Inc.

package param

// GetHostPhysicalMemoryFactsDetailParam GetHostPhysicalMemoryFacts详细参数
type GetHostPhysicalMemoryFactsDetailParam struct {
	rest string `json:"hostUuid" validate:"required"` // 必填
}

// GetHostPhysicalMemoryFactsParam GetHostPhysicalMemoryFacts请求参数
type GetHostPhysicalMemoryFactsParam struct {
	BaseParam
	Params GetHostPhysicalMemoryFactsDetailParam `json:"params"` // 详细参数
}

