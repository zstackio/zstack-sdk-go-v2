// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVRouterDistributedRoutingEnabledDetailParam GetVpcVRouterDistributedRoutingEnabled详细参数
type GetVpcVRouterDistributedRoutingEnabledDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVpcVRouterDistributedRoutingEnabledParam GetVpcVRouterDistributedRoutingEnabled请求参数
type GetVpcVRouterDistributedRoutingEnabledParam struct {
	BaseParam
	Params GetVpcVRouterDistributedRoutingEnabledDetailParam `json:"params"` // 详细参数
}

