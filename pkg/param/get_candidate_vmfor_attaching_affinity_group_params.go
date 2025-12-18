// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVMForAttachingAffinityGroupDetailParam GetCandidateVMForAttachingAffinityGroup detail param
type GetCandidateVMForAttachingAffinityGroupDetailParam struct {
	AffinityGroupUuid string `json:"affinityGroupUuid" validate:"required"`
}

// GetCandidateVMForAttachingAffinityGroupParam GetCandidateVMForAttachingAffinityGroup request param
type GetCandidateVMForAttachingAffinityGroupParam struct {
	BaseParam
	Params GetCandidateVMForAttachingAffinityGroupDetailParam `json:"params"`
}
