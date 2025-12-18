// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForLoadBalancerDetailParam GetCandidateL3NetworksForLoadBalancer详细参数
type GetCandidateL3NetworksForLoadBalancerDetailParam struct {
	rest string `json:"listenerUuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForLoadBalancerParam GetCandidateL3NetworksForLoadBalancer请求参数
type GetCandidateL3NetworksForLoadBalancerParam struct {
	BaseParam
	Params GetCandidateL3NetworksForLoadBalancerDetailParam `json:"params"` // 详细参数
}

