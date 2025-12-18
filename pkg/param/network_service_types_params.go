// Copyright (c) ZStack.io, Inc.

package param

// GetNetworkServiceTypesDetailParam GetNetworkServiceTypes详细参数
type GetNetworkServiceTypesDetailParam struct {
}

// GetNetworkServiceTypesParam GetNetworkServiceTypes请求参数
type GetNetworkServiceTypesParam struct {
	BaseParam
	Params GetNetworkServiceTypesDetailParam `json:"params"` // 详细参数
}

