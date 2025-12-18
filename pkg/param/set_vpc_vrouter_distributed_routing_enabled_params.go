// Copyright (c) ZStack.io, Inc.

package param

// SetVpcVRouterDistributedRoutingEnabledDetailParam SetVpcVRouterDistributedRoutingEnabled detail param
type SetVpcVRouterDistributedRoutingEnabledDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// SetVpcVRouterDistributedRoutingEnabledParam SetVpcVRouterDistributedRoutingEnabled request param
type SetVpcVRouterDistributedRoutingEnabledParam struct {
	BaseParam
	Params SetVpcVRouterDistributedRoutingEnabledDetailParam `json:"params"`
}
