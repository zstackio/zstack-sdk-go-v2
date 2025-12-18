// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmForAttachingIsoDetailParam GetCandidateVmForAttachingIso详细参数
type GetCandidateVmForAttachingIsoDetailParam struct {
	rest string `json:"isoUuid" validate:"required"` // 必填
}

// GetCandidateVmForAttachingIsoParam GetCandidateVmForAttachingIso请求参数
type GetCandidateVmForAttachingIsoParam struct {
	BaseParam
	Params GetCandidateVmForAttachingIsoDetailParam `json:"params"` // 详细参数
}

