// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForServerGroupDetailParam GetCandidateL3NetworksForServerGroup详细参数
type GetCandidateL3NetworksForServerGroupDetailParam struct {
	rest string `json:"serverGroupUuid,omitempty"`
	rest string `json:"loadBalancerUuid,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForServerGroupParam GetCandidateL3NetworksForServerGroup请求参数
type GetCandidateL3NetworksForServerGroupParam struct {
	BaseParam
	Params GetCandidateL3NetworksForServerGroupDetailParam `json:"params"` // 详细参数
}

