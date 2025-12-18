// Copyright (c) ZStack.io, Inc.

package param

// DeleteLoadBalancerServerGroupDetailParam DeleteLoadBalancerServerGroup detail param
type DeleteLoadBalancerServerGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteLoadBalancerServerGroupParam DeleteLoadBalancerServerGroup request param
type DeleteLoadBalancerServerGroupParam struct {
	BaseParam
	Params DeleteLoadBalancerServerGroupDetailParam `json:"params"`
}
