// Copyright (c) ZStack.io, Inc.

package param

// AddSecurityGroupRuleDetailParam AddSecurityGroupRule detail param
type AddSecurityGroupRuleDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	Rules []SecurityGroupRuleAOParam `json:"rules" validate:"required"`
	RemoteSecurityGroupUuids []string `json:"remoteSecurityGroupUuids,omitempty"`
	Priority int `json:"priority,omitempty"`
}

// AddSecurityGroupRuleParam AddSecurityGroupRule request param
type AddSecurityGroupRuleParam struct {
	BaseParam
	Params AddSecurityGroupRuleDetailParam `json:"params"`
}
