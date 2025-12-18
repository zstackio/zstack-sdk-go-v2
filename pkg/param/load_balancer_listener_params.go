// Copyright (c) ZStack.io, Inc.

package param

// UpdateLoadBalancerListenerDetailParam UpdateLoadBalancerListener详细参数
type UpdateLoadBalancerListenerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateLoadBalancerListenerParam UpdateLoadBalancerListener请求参数
type UpdateLoadBalancerListenerParam struct {
	BaseParam
	Params UpdateLoadBalancerListenerDetailParam `json:"params"` // 详细参数
}

