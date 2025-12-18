// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL2NetworksForAttachingClusterDetailParam GetCandidateL2NetworksForAttachingCluster详细参数
type GetCandidateL2NetworksForAttachingClusterDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCandidateL2NetworksForAttachingClusterParam GetCandidateL2NetworksForAttachingCluster请求参数
type GetCandidateL2NetworksForAttachingClusterParam struct {
	BaseParam
	Params GetCandidateL2NetworksForAttachingClusterDetailParam `json:"params"` // 详细参数
}

