// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectContainerClusterCandidatesDetailParam GetIAM2ProjectContainerClusterCandidates详细参数
type GetIAM2ProjectContainerClusterCandidatesDetailParam struct {
}

// GetIAM2ProjectContainerClusterCandidatesParam GetIAM2ProjectContainerClusterCandidates请求参数
type GetIAM2ProjectContainerClusterCandidatesParam struct {
	BaseParam
	Params GetIAM2ProjectContainerClusterCandidatesDetailParam `json:"params"` // 详细参数
}

