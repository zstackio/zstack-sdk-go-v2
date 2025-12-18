// Copyright (c) ZStack.io, Inc.

package param

// GetSharedBlockCandidateDetailParam GetSharedBlockCandidate详细参数
type GetSharedBlockCandidateDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// GetSharedBlockCandidateParam GetSharedBlockCandidate请求参数
type GetSharedBlockCandidateParam struct {
	BaseParam
	Params GetSharedBlockCandidateDetailParam `json:"params"` // 详细参数
}

