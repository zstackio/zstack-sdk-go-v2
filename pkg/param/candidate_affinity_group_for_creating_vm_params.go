// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateAffinityGroupForCreatingVmDetailParam GetCandidateAffinityGroupForCreatingVm详细参数
type GetCandidateAffinityGroupForCreatingVmDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
}

// GetCandidateAffinityGroupForCreatingVmParam GetCandidateAffinityGroupForCreatingVm请求参数
type GetCandidateAffinityGroupForCreatingVmParam struct {
	BaseParam
	Params GetCandidateAffinityGroupForCreatingVmDetailParam `json:"params"` // 详细参数
}

