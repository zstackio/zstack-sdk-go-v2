// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateAffinityGroupForCreatingVmDetailParam GetCandidateAffinityGroupForCreatingVm detail param
type GetCandidateAffinityGroupForCreatingVmDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
}

// GetCandidateAffinityGroupForCreatingVmParam GetCandidateAffinityGroupForCreatingVm request param
type GetCandidateAffinityGroupForCreatingVmParam struct {
	BaseParam
	Params GetCandidateAffinityGroupForCreatingVmDetailParam `json:"params"`
}
