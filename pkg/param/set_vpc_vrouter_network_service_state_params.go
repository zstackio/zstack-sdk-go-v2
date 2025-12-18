// Copyright (c) ZStack.io, Inc.

package param

// SetVpcVRouterNetworkServiceStateDetailParam SetVpcVRouterNetworkServiceState detail param
type SetVpcVRouterNetworkServiceStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	NetworkService string `json:"networkService" validate:"required"`
	State string `json:"state" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
}

// SetVpcVRouterNetworkServiceStateParam SetVpcVRouterNetworkServiceState request param
type SetVpcVRouterNetworkServiceStateParam struct {
	BaseParam
	Params SetVpcVRouterNetworkServiceStateDetailParam `json:"params"`
}
