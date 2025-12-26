// Copyright (c) ZStack.io, Inc.

package param

// UpdateSecurityGroupRulePriorityDetailParam UpdateSecurityGroupRulePriority detail param
type UpdateSecurityGroupRulePriorityDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	Rules []SecurityGroupRulePriorityAOParam `json:"rules" validate:"required"`
}

// UpdateSecurityGroupRulePriorityParam UpdateSecurityGroupRulePriority request param
type UpdateSecurityGroupRulePriorityParam struct {
	BaseParam
	Params UpdateSecurityGroupRulePriorityDetailParam `json:"params"`
}
