// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateIsoForAttachingVmDetailParam GetCandidateIsoForAttachingVm详细参数
type GetCandidateIsoForAttachingVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetCandidateIsoForAttachingVmParam GetCandidateIsoForAttachingVm请求参数
type GetCandidateIsoForAttachingVmParam struct {
	BaseParam
	Params GetCandidateIsoForAttachingVmDetailParam `json:"params"` // 详细参数
}

