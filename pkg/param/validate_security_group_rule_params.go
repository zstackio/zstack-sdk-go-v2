// Copyright (c) ZStack.io, Inc.

package param

// ValidateSecurityGroupRuleDetailParam ValidateSecurityGroupRule详细参数
type ValidateSecurityGroupRuleDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"protocol" validate:"required"` // 必填
	rest string `json:"remoteSecurityGroupUuid,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"srcIpRange,omitempty"`
	rest string `json:"dstIpRange,omitempty"`
	rest string `json:"dstPortRange,omitempty"`
	rest string `json:"action,omitempty"`
	rest int `json:"startPort,omitempty"`
	rest int `json:"endPort,omitempty"`
	rest string `json:"allowedCidr,omitempty"`
}

// ValidateSecurityGroupRuleParam ValidateSecurityGroupRule请求参数
type ValidateSecurityGroupRuleParam struct {
	BaseParam
	Params ValidateSecurityGroupRuleDetailParam `json:"params"` // 详细参数
}

