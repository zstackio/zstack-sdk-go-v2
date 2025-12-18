// Copyright (c) ZStack.io, Inc.

package param

// UpdateLoadBalancerListenerDetailParam UpdateLoadBalancerListener detail param
type UpdateLoadBalancerListenerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateLoadBalancerListenerParam UpdateLoadBalancerListener request param
type UpdateLoadBalancerListenerParam struct {
	BaseParam
	Params UpdateLoadBalancerListenerDetailParam `json:"params"`
}
