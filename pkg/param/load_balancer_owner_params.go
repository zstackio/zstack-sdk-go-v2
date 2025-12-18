// Copyright (c) ZStack.io, Inc.

package param

// GetLoadBalancerOwnerDetailParam GetLoadBalancerOwner详细参数
type GetLoadBalancerOwnerDetailParam struct {
	rest string `json:"loadBalancerUuid" validate:"required"` // 必填
}

// GetLoadBalancerOwnerParam GetLoadBalancerOwner请求参数
type GetLoadBalancerOwnerParam struct {
	BaseParam
	Params GetLoadBalancerOwnerDetailParam `json:"params"` // 详细参数
}

