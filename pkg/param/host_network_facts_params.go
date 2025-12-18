// Copyright (c) ZStack.io, Inc.

package param

// GetHostNetworkFactsDetailParam GetHostNetworkFacts详细参数
type GetHostNetworkFactsDetailParam struct {
	rest string `json:"hostUuid" validate:"required"` // 必填
}

// GetHostNetworkFactsParam GetHostNetworkFacts请求参数
type GetHostNetworkFactsParam struct {
	BaseParam
	Params GetHostNetworkFactsDetailParam `json:"params"` // 详细参数
}

