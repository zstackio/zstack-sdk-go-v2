// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmForAttachingIsoDetailParam GetCandidateVmForAttachingIso detail param
type GetCandidateVmForAttachingIsoDetailParam struct {
	IsoUuid string `json:"isoUuid" validate:"required"`
}

// GetCandidateVmForAttachingIsoParam GetCandidateVmForAttachingIso request param
type GetCandidateVmForAttachingIsoParam struct {
	BaseParam
	Params GetCandidateVmForAttachingIsoDetailParam `json:"params"`
}
