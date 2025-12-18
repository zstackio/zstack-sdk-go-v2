// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecurityGroupRuleDetailParam ChangeSecurityGroupRule detail param
type ChangeSecurityGroupRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	RemoteSecurityGroupUuid string `json:"remoteSecurityGroupUuid,omitempty"`
	Action string `json:"action,omitempty"`
	State string `json:"state,omitempty"`
	Priority int `json:"priority,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	SrcIpRange string `json:"srcIpRange,omitempty"`
	DstIpRange string `json:"dstIpRange,omitempty"`
	DstPortRange string `json:"dstPortRange,omitempty"`
}

// ChangeSecurityGroupRuleParam ChangeSecurityGroupRule request param
type ChangeSecurityGroupRuleParam struct {
	BaseParam
	Params ChangeSecurityGroupRuleDetailParam `json:"params"`
}
