// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVRouterNetworkServiceStateDetailParam GetVpcVRouterNetworkServiceState detail param
type GetVpcVRouterNetworkServiceStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	NetworkService string `json:"networkService" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
}

// GetVpcVRouterNetworkServiceStateParam GetVpcVRouterNetworkServiceState request param
type GetVpcVRouterNetworkServiceStateParam struct {
	BaseParam
	Params GetVpcVRouterNetworkServiceStateDetailParam `json:"params"`
}
