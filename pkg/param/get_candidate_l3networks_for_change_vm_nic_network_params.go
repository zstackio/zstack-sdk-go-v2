// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForChangeVmNicNetworkDetailParam GetCandidateL3NetworksForChangeVmNicNetwork detail param
type GetCandidateL3NetworksForChangeVmNicNetworkDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
}

// GetCandidateL3NetworksForChangeVmNicNetworkParam GetCandidateL3NetworksForChangeVmNicNetwork request param
type GetCandidateL3NetworksForChangeVmNicNetworkParam struct {
	BaseParam
	Params GetCandidateL3NetworksForChangeVmNicNetworkDetailParam `json:"params"`
}
