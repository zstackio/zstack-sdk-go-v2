// Copyright (c) ZStack.io, Inc.

package param

// GetClusterHostNetworkFactsDetailParam GetClusterHostNetworkFacts详细参数
type GetClusterHostNetworkFactsDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetClusterHostNetworkFactsParam GetClusterHostNetworkFacts请求参数
type GetClusterHostNetworkFactsParam struct {
	BaseParam
	Params GetClusterHostNetworkFactsDetailParam `json:"params"` // 详细参数
}

