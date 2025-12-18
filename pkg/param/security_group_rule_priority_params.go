// Copyright (c) ZStack.io, Inc.

package param

// UpdateSecurityGroupRulePriorityDetailParam UpdateSecurityGroupRulePriority详细参数
type UpdateSecurityGroupRulePriorityDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest []interface{} `json:"rules" validate:"required"` // 必填
}

// UpdateSecurityGroupRulePriorityParam UpdateSecurityGroupRulePriority请求参数
type UpdateSecurityGroupRulePriorityParam struct {
	BaseParam
	Params UpdateSecurityGroupRulePriorityDetailParam `json:"params"` // 详细参数
}

