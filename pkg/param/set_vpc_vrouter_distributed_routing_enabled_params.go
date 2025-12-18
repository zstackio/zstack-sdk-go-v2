// Copyright (c) ZStack.io, Inc.

package param

// SetVpcVRouterDistributedRoutingEnabledDetailParam SetVpcVRouterDistributedRoutingEnabled详细参数
type SetVpcVRouterDistributedRoutingEnabledDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// SetVpcVRouterDistributedRoutingEnabledParam SetVpcVRouterDistributedRoutingEnabled请求参数
type SetVpcVRouterDistributedRoutingEnabledParam struct {
	BaseParam
	Params SetVpcVRouterDistributedRoutingEnabledDetailParam `json:"params"` // 详细参数
}

