// Copyright (c) ZStack.io, Inc.

package param

// GetClusterHostNetworkFactsDetailParam GetClusterHostNetworkFacts detail param
type GetClusterHostNetworkFactsDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetClusterHostNetworkFactsParam GetClusterHostNetworkFacts request param
type GetClusterHostNetworkFactsParam struct {
	BaseParam
	Params GetClusterHostNetworkFactsDetailParam `json:"params"`
}
