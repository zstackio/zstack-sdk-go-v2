// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicsForLoadBalancerServerGroupDetailParam GetCandidateVmNicsForLoadBalancerServerGroup detail param
type GetCandidateVmNicsForLoadBalancerServerGroupDetailParam struct {
	ServergroupUuid string `json:"servergroupUuid,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
}

// GetCandidateVmNicsForLoadBalancerServerGroupParam GetCandidateVmNicsForLoadBalancerServerGroup request param
type GetCandidateVmNicsForLoadBalancerServerGroupParam struct {
	BaseParam
	Params GetCandidateVmNicsForLoadBalancerServerGroupDetailParam `json:"params"`
}
