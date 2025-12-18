// Copyright (c) ZStack.io, Inc.

package param

// RefreshLoadBalancerDetailParam RefreshLoadBalancer detail param
type RefreshLoadBalancerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RefreshLoadBalancerParam RefreshLoadBalancer request param
type RefreshLoadBalancerParam struct {
	BaseParam
	Params RefreshLoadBalancerDetailParam `json:"params"`
}
