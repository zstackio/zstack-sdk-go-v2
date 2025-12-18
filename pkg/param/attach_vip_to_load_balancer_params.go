// Copyright (c) ZStack.io, Inc.

package param

// AttachVipToLoadBalancerDetailParam AttachVipToLoadBalancer detail param
type AttachVipToLoadBalancerDetailParam struct {
	LoadBalancerUuid string `json:"loadBalancerUuid" validate:"required"`
	VipUuid string `json:"vipUuid" validate:"required"`
}

// AttachVipToLoadBalancerParam AttachVipToLoadBalancer request param
type AttachVipToLoadBalancerParam struct {
	BaseParam
	Params AttachVipToLoadBalancerDetailParam `json:"params"`
}
