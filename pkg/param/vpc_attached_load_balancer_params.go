// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedLoadBalancerDetailParam GetVpcAttachedLoadBalancer详细参数
type GetVpcAttachedLoadBalancerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVpcAttachedLoadBalancerParam GetVpcAttachedLoadBalancer请求参数
type GetVpcAttachedLoadBalancerParam struct {
	BaseParam
	Params GetVpcAttachedLoadBalancerDetailParam `json:"params"` // 详细参数
}

