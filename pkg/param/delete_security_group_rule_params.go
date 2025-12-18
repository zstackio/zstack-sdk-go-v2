// Copyright (c) ZStack.io, Inc.

package param

// DeleteSecurityGroupRuleDetailParam DeleteSecurityGroupRule detail param
type DeleteSecurityGroupRuleDetailParam struct {
	RuleUuids []string `json:"ruleUuids" validate:"required"`
}

// DeleteSecurityGroupRuleParam DeleteSecurityGroupRule request param
type DeleteSecurityGroupRuleParam struct {
	BaseParam
	Params DeleteSecurityGroupRuleDetailParam `json:"params"`
}
