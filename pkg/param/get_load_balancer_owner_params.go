// Copyright (c) ZStack.io, Inc.

package param

// GetLoadBalancerOwnerDetailParam GetLoadBalancerOwner detail param
type GetLoadBalancerOwnerDetailParam struct {
	LoadBalancerUuid string `json:"loadBalancerUuid" validate:"required"`
}

// GetLoadBalancerOwnerParam GetLoadBalancerOwner request param
type GetLoadBalancerOwnerParam struct {
	BaseParam
	Params GetLoadBalancerOwnerDetailParam `json:"params"`
}
