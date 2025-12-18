// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVMForAttachingAffinityGroupDetailParam GetCandidateVMForAttachingAffinityGroup详细参数
type GetCandidateVMForAttachingAffinityGroupDetailParam struct {
	rest string `json:"affinityGroupUuid" validate:"required"` // 必填
}

// GetCandidateVMForAttachingAffinityGroupParam GetCandidateVMForAttachingAffinityGroup请求参数
type GetCandidateVMForAttachingAffinityGroupParam struct {
	BaseParam
	Params GetCandidateVMForAttachingAffinityGroupDetailParam `json:"params"` // 详细参数
}

