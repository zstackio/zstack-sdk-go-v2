// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForServerGroupDetailParam GetCandidateL3NetworksForServerGroup detail param
type GetCandidateL3NetworksForServerGroupDetailParam struct {
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForServerGroupParam GetCandidateL3NetworksForServerGroup request param
type GetCandidateL3NetworksForServerGroupParam struct {
	BaseParam
	Params GetCandidateL3NetworksForServerGroupDetailParam `json:"params"`
}
