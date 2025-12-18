// Copyright (c) ZStack.io, Inc.

package param

// GetHypervisorTypesDetailParam GetHypervisorTypes详细参数
type GetHypervisorTypesDetailParam struct {
}

// GetHypervisorTypesParam GetHypervisorTypes请求参数
type GetHypervisorTypesParam struct {
	BaseParam
	Params GetHypervisorTypesDetailParam `json:"params"` // 详细参数
}

