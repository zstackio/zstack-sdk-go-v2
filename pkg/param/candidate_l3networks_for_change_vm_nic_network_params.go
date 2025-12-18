// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForChangeVmNicNetworkDetailParam GetCandidateL3NetworksForChangeVmNicNetwork详细参数
type GetCandidateL3NetworksForChangeVmNicNetworkDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
}

// GetCandidateL3NetworksForChangeVmNicNetworkParam GetCandidateL3NetworksForChangeVmNicNetwork请求参数
type GetCandidateL3NetworksForChangeVmNicNetworkParam struct {
	BaseParam
	Params GetCandidateL3NetworksForChangeVmNicNetworkDetailParam `json:"params"` // 详细参数
}

