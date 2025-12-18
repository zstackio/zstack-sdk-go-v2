// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicForSecurityGroupDetailParam GetCandidateVmNicForSecurityGroup detail param
type GetCandidateVmNicForSecurityGroupDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
}

// GetCandidateVmNicForSecurityGroupParam GetCandidateVmNicForSecurityGroup request param
type GetCandidateVmNicForSecurityGroupParam struct {
	BaseParam
	Params GetCandidateVmNicForSecurityGroupDetailParam `json:"params"`
}
