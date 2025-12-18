// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicsForLoadBalancerDetailParam GetCandidateVmNicsForLoadBalancer detail param
type GetCandidateVmNicsForLoadBalancerDetailParam struct {
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// GetCandidateVmNicsForLoadBalancerParam GetCandidateVmNicsForLoadBalancer request param
type GetCandidateVmNicsForLoadBalancerParam struct {
	BaseParam
	Params GetCandidateVmNicsForLoadBalancerDetailParam `json:"params"`
}
