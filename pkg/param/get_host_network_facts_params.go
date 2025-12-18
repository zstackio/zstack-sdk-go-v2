// Copyright (c) ZStack.io, Inc.

package param

// GetHostNetworkFactsDetailParam GetHostNetworkFacts detail param
type GetHostNetworkFactsDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
}

// GetHostNetworkFactsParam GetHostNetworkFacts request param
type GetHostNetworkFactsParam struct {
	BaseParam
	Params GetHostNetworkFactsDetailParam `json:"params"`
}
