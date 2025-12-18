// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateAffinityGroupForAttachingVmDetailParam GetCandidateAffinityGroupForAttachingVm详细参数
type GetCandidateAffinityGroupForAttachingVmDetailParam struct {
	rest string `json:"vmUuid" validate:"required"` // 必填
}

// GetCandidateAffinityGroupForAttachingVmParam GetCandidateAffinityGroupForAttachingVm请求参数
type GetCandidateAffinityGroupForAttachingVmParam struct {
	BaseParam
	Params GetCandidateAffinityGroupForAttachingVmDetailParam `json:"params"` // 详细参数
}

