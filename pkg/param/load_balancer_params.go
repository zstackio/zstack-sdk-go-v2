// Copyright (c) ZStack.io, Inc.

package param

// DeleteLoadBalancerDetailParam DeleteLoadBalancer详细参数
type DeleteLoadBalancerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteLoadBalancerParam DeleteLoadBalancer请求参数
type DeleteLoadBalancerParam struct {
	BaseParam
	Params DeleteLoadBalancerDetailParam `json:"params"` // 详细参数
}

