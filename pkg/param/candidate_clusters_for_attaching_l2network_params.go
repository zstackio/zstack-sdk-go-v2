// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateClustersForAttachingL2NetworkDetailParam GetCandidateClustersForAttachingL2Network详细参数
type GetCandidateClustersForAttachingL2NetworkDetailParam struct {
	rest string `json:"l2NetworkUuid" validate:"required"` // 必填
	rest []string `json:"clusterTypes,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCandidateClustersForAttachingL2NetworkParam GetCandidateClustersForAttachingL2Network请求参数
type GetCandidateClustersForAttachingL2NetworkParam struct {
	BaseParam
	Params GetCandidateClustersForAttachingL2NetworkDetailParam `json:"params"` // 详细参数
}

