// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForLoadBalancerDetailParam GetCandidateL3NetworksForLoadBalancer detail param
type GetCandidateL3NetworksForLoadBalancerDetailParam struct {
	ListenerUuid string `json:"listenerUuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForLoadBalancerParam GetCandidateL3NetworksForLoadBalancer request param
type GetCandidateL3NetworksForLoadBalancerParam struct {
	BaseParam
	Params GetCandidateL3NetworksForLoadBalancerDetailParam `json:"params"`
}
