// Copyright (c) ZStack.io, Inc.

package param

// DeleteLoadBalancerListenerDetailParam DeleteLoadBalancerListener detail param
type DeleteLoadBalancerListenerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteLoadBalancerListenerParam DeleteLoadBalancerListener request param
type DeleteLoadBalancerListenerParam struct {
	BaseParam
	Params DeleteLoadBalancerListenerDetailParam `json:"params"`
}
