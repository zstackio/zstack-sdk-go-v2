// Copyright (c) ZStack.io, Inc.

package param

// AttachVipToLoadBalancerDetailParam AttachVipToLoadBalancer详细参数
type AttachVipToLoadBalancerDetailParam struct {
	rest string `json:"loadBalancerUuid" validate:"required"` // 必填
	rest string `json:"vipUuid" validate:"required"` // 必填
}

// AttachVipToLoadBalancerParam AttachVipToLoadBalancer请求参数
type AttachVipToLoadBalancerParam struct {
	BaseParam
	Params AttachVipToLoadBalancerDetailParam `json:"params"` // 详细参数
}

