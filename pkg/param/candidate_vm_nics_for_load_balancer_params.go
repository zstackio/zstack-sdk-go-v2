// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicsForLoadBalancerDetailParam GetCandidateVmNicsForLoadBalancer详细参数
type GetCandidateVmNicsForLoadBalancerDetailParam struct {
	rest string `json:"listenerUuid" validate:"required"` // 必填
}

// GetCandidateVmNicsForLoadBalancerParam GetCandidateVmNicsForLoadBalancer请求参数
type GetCandidateVmNicsForLoadBalancerParam struct {
	BaseParam
	Params GetCandidateVmNicsForLoadBalancerDetailParam `json:"params"` // 详细参数
}

