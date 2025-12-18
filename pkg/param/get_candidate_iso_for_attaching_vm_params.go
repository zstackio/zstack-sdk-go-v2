// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateIsoForAttachingVmDetailParam GetCandidateIsoForAttachingVm detail param
type GetCandidateIsoForAttachingVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetCandidateIsoForAttachingVmParam GetCandidateIsoForAttachingVm request param
type GetCandidateIsoForAttachingVmParam struct {
	BaseParam
	Params GetCandidateIsoForAttachingVmDetailParam `json:"params"`
}
