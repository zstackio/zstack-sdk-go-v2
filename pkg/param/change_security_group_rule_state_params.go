// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecurityGroupRuleStateDetailParam ChangeSecurityGroupRuleState detail param
type ChangeSecurityGroupRuleStateDetailParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid" validate:"required"`
	RuleUuids []string `json:"ruleUuids" validate:"required"`
	State string `json:"state" validate:"required"`
}

// ChangeSecurityGroupRuleStateParam ChangeSecurityGroupRuleState request param
type ChangeSecurityGroupRuleStateParam struct {
	BaseParam
	Params ChangeSecurityGroupRuleStateDetailParam `json:"params"`
}
