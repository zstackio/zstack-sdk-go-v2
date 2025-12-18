// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVRouterNetworkServiceStateDetailParam GetVpcVRouterNetworkServiceState详细参数
type GetVpcVRouterNetworkServiceStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"networkService" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid,omitempty"`
}

// GetVpcVRouterNetworkServiceStateParam GetVpcVRouterNetworkServiceState请求参数
type GetVpcVRouterNetworkServiceStateParam struct {
	BaseParam
	Params GetVpcVRouterNetworkServiceStateDetailParam `json:"params"` // 详细参数
}

