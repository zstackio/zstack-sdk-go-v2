// Copyright (c) ZStack.io, Inc.

package param

// ValidateSecurityGroupRuleDetailParam ValidateSecurityGroupRule detail param
type ValidateSecurityGroupRuleDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	Protocol string `json:"protocol" validate:"required"`
	RemoteSecurityGroupUuid string `json:"remoteSecurityGroupUuid,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	SrcIpRange string `json:"srcIpRange,omitempty"`
	DstIpRange string `json:"dstIpRange,omitempty"`
	DstPortRange string `json:"dstPortRange,omitempty"`
	Action string `json:"action,omitempty"`
	StartPort int `json:"startPort,omitempty"`
	EndPort int `json:"endPort,omitempty"`
	AllowedCidr string `json:"allowedCidr,omitempty"`
}

// ValidateSecurityGroupRuleParam ValidateSecurityGroupRule request param
type ValidateSecurityGroupRuleParam struct {
	BaseParam
	Params ValidateSecurityGroupRuleDetailParam `json:"params"`
}
