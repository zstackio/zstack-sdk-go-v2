// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedLoadBalancerDetailParam GetVpcAttachedLoadBalancer detail param
type GetVpcAttachedLoadBalancerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVpcAttachedLoadBalancerParam GetVpcAttachedLoadBalancer request param
type GetVpcAttachedLoadBalancerParam struct {
	BaseParam
	Params GetVpcAttachedLoadBalancerDetailParam `json:"params"`
}
