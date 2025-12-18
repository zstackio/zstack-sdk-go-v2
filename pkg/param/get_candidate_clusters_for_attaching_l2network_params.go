// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateClustersForAttachingL2NetworkDetailParam GetCandidateClustersForAttachingL2Network detail param
type GetCandidateClustersForAttachingL2NetworkDetailParam struct {
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	ClusterTypes []string `json:"clusterTypes,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCandidateClustersForAttachingL2NetworkParam GetCandidateClustersForAttachingL2Network request param
type GetCandidateClustersForAttachingL2NetworkParam struct {
	BaseParam
	Params GetCandidateClustersForAttachingL2NetworkDetailParam `json:"params"`
}
