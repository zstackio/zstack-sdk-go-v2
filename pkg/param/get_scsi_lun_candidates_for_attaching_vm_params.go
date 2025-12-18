// Copyright (c) ZStack.io, Inc.

package param

// GetScsiLunCandidatesForAttachingVmDetailParam GetScsiLunCandidatesForAttachingVm detail param
type GetScsiLunCandidatesForAttachingVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetScsiLunCandidatesForAttachingVmParam GetScsiLunCandidatesForAttachingVm request param
type GetScsiLunCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetScsiLunCandidatesForAttachingVmDetailParam `json:"params"`
}
