// Copyright (c) ZStack.io, Inc.

package param

// GetSharedBlockCandidateDetailParam GetSharedBlockCandidate detail param
type GetSharedBlockCandidateDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// GetSharedBlockCandidateParam GetSharedBlockCandidate request param
type GetSharedBlockCandidateParam struct {
	BaseParam
	Params GetSharedBlockCandidateDetailParam `json:"params"`
}
