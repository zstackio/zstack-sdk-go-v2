// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicForSecurityGroupDetailParam GetCandidateVmNicForSecurityGroup详细参数
type GetCandidateVmNicForSecurityGroupDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
}

// GetCandidateVmNicForSecurityGroupParam GetCandidateVmNicForSecurityGroup请求参数
type GetCandidateVmNicForSecurityGroupParam struct {
	BaseParam
	Params GetCandidateVmNicForSecurityGroupDetailParam `json:"params"` // 详细参数
}

