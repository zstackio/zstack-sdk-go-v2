// Copyright (c) ZStack.io, Inc.

package param

// GetContainerUsageDetailParam GetContainerUsage detail param
type GetContainerUsageDetailParam struct {
}

// GetContainerUsageParam GetContainerUsage request param
type GetContainerUsageParam struct {
	BaseParam
	Params GetContainerUsageDetailParam `json:"params"`
}
