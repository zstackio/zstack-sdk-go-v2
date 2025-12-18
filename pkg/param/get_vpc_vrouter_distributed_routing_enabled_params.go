// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVRouterDistributedRoutingEnabledDetailParam GetVpcVRouterDistributedRoutingEnabled detail param
type GetVpcVRouterDistributedRoutingEnabledDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVpcVRouterDistributedRoutingEnabledParam GetVpcVRouterDistributedRoutingEnabled request param
type GetVpcVRouterDistributedRoutingEnabledParam struct {
	BaseParam
	Params GetVpcVRouterDistributedRoutingEnabledDetailParam `json:"params"`
}
