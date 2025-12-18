// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecurityGroupRuleStateDetailParam ChangeSecurityGroupRuleState详细参数
type ChangeSecurityGroupRuleStateDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest []string `json:"ruleUuids" validate:"required"` // 必填
	rest string `json:"state" validate:"required"` // 必填
}

// ChangeSecurityGroupRuleStateParam ChangeSecurityGroupRuleState请求参数
type ChangeSecurityGroupRuleStateParam struct {
	BaseParam
	Params ChangeSecurityGroupRuleStateDetailParam `json:"params"` // 详细参数
}

