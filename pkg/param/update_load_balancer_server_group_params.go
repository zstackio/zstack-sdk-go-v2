// Copyright (c) ZStack.io, Inc.

package param

// UpdateLoadBalancerServerGroupDetailParam UpdateLoadBalancerServerGroup detail param
type UpdateLoadBalancerServerGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateLoadBalancerServerGroupParam UpdateLoadBalancerServerGroup request param
type UpdateLoadBalancerServerGroupParam struct {
	BaseParam
	Params UpdateLoadBalancerServerGroupDetailParam `json:"params"`
}
