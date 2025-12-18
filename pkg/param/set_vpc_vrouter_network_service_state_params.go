// Copyright (c) ZStack.io, Inc.

package param

// SetVpcVRouterNetworkServiceStateDetailParam SetVpcVRouterNetworkServiceState详细参数
type SetVpcVRouterNetworkServiceStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"networkService" validate:"required"` // 必填
	rest string `json:"state" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid,omitempty"`
}

// SetVpcVRouterNetworkServiceStateParam SetVpcVRouterNetworkServiceState请求参数
type SetVpcVRouterNetworkServiceStateParam struct {
	BaseParam
	Params SetVpcVRouterNetworkServiceStateDetailParam `json:"params"` // 详细参数
}

