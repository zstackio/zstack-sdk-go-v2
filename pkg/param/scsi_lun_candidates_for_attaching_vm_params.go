// Copyright (c) ZStack.io, Inc.

package param

// GetScsiLunCandidatesForAttachingVmDetailParam GetScsiLunCandidatesForAttachingVm详细参数
type GetScsiLunCandidatesForAttachingVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetScsiLunCandidatesForAttachingVmParam GetScsiLunCandidatesForAttachingVm请求参数
type GetScsiLunCandidatesForAttachingVmParam struct {
	BaseParam
	Params GetScsiLunCandidatesForAttachingVmDetailParam `json:"params"` // 详细参数
}

