// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicsForLoadBalancerServerGroupDetailParam GetCandidateVmNicsForLoadBalancerServerGroup详细参数
type GetCandidateVmNicsForLoadBalancerServerGroupDetailParam struct {
	rest string `json:"servergroupUuid,omitempty"`
	rest string `json:"loadBalancerUuid,omitempty"`
	rest int `json:"ipVersion,omitempty"`
}

// GetCandidateVmNicsForLoadBalancerServerGroupParam GetCandidateVmNicsForLoadBalancerServerGroup请求参数
type GetCandidateVmNicsForLoadBalancerServerGroupParam struct {
	BaseParam
	Params GetCandidateVmNicsForLoadBalancerServerGroupDetailParam `json:"params"` // 详细参数
}

