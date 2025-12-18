// Copyright (c) ZStack.io, Inc.

package param

// RefreshLoadBalancerDetailParam RefreshLoadBalancer详细参数
type RefreshLoadBalancerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RefreshLoadBalancerParam RefreshLoadBalancer请求参数
type RefreshLoadBalancerParam struct {
	BaseParam
	Params RefreshLoadBalancerDetailParam `json:"params"` // 详细参数
}

