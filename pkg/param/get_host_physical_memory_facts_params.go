// Copyright (c) ZStack.io, Inc.

package param

// GetHostPhysicalMemoryFactsDetailParam GetHostPhysicalMemoryFacts detail param
type GetHostPhysicalMemoryFactsDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
}

// GetHostPhysicalMemoryFactsParam GetHostPhysicalMemoryFacts request param
type GetHostPhysicalMemoryFactsParam struct {
	BaseParam
	Params GetHostPhysicalMemoryFactsDetailParam `json:"params"`
}
