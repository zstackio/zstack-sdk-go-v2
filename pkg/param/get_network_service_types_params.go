// Copyright (c) ZStack.io, Inc.

package param

// GetNetworkServiceTypesDetailParam GetNetworkServiceTypes detail param
type GetNetworkServiceTypesDetailParam struct {
}

// GetNetworkServiceTypesParam GetNetworkServiceTypes request param
type GetNetworkServiceTypesParam struct {
	BaseParam
	Params GetNetworkServiceTypesDetailParam `json:"params"`
}
