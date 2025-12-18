// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicsForPortMirrorDetailParam GetCandidateVmNicsForPortMirror详细参数
type GetCandidateVmNicsForPortMirrorDetailParam struct {
	rest string `json:"portMirrorUuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
}

// GetCandidateVmNicsForPortMirrorParam GetCandidateVmNicsForPortMirror请求参数
type GetCandidateVmNicsForPortMirrorParam struct {
	BaseParam
	Params GetCandidateVmNicsForPortMirrorDetailParam `json:"params"` // 详细参数
}

