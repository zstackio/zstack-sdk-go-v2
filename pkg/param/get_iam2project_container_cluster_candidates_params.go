// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectContainerClusterCandidatesDetailParam GetIAM2ProjectContainerClusterCandidates detail param
type GetIAM2ProjectContainerClusterCandidatesDetailParam struct {
}

// GetIAM2ProjectContainerClusterCandidatesParam GetIAM2ProjectContainerClusterCandidates request param
type GetIAM2ProjectContainerClusterCandidatesParam struct {
	BaseParam
	Params GetIAM2ProjectContainerClusterCandidatesDetailParam `json:"params"`
}
