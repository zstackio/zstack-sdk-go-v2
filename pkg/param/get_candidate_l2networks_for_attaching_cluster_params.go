// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL2NetworksForAttachingClusterDetailParam GetCandidateL2NetworksForAttachingCluster detail param
type GetCandidateL2NetworksForAttachingClusterDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCandidateL2NetworksForAttachingClusterParam GetCandidateL2NetworksForAttachingCluster request param
type GetCandidateL2NetworksForAttachingClusterParam struct {
	BaseParam
	Params GetCandidateL2NetworksForAttachingClusterDetailParam `json:"params"`
}
