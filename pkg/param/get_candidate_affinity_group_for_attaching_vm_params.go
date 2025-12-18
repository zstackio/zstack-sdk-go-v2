// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateAffinityGroupForAttachingVmDetailParam GetCandidateAffinityGroupForAttachingVm detail param
type GetCandidateAffinityGroupForAttachingVmDetailParam struct {
	VmUuid string `json:"vmUuid" validate:"required"`
}

// GetCandidateAffinityGroupForAttachingVmParam GetCandidateAffinityGroupForAttachingVm request param
type GetCandidateAffinityGroupForAttachingVmParam struct {
	BaseParam
	Params GetCandidateAffinityGroupForAttachingVmDetailParam `json:"params"`
}
