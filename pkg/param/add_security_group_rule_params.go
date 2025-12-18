// Copyright (c) ZStack.io, Inc.

package param

// AddSecurityGroupRuleDetailParam AddSecurityGroupRule详细参数
type AddSecurityGroupRuleDetailParam struct {
	rest string `json:"securityGroupUuid" validate:"required"` // 必填
	rest []interface{} `json:"rules" validate:"required"` // 必填
	rest []string `json:"remoteSecurityGroupUuids,omitempty"`
	rest int `json:"priority,omitempty"`
}

// AddSecurityGroupRuleParam AddSecurityGroupRule请求参数
type AddSecurityGroupRuleParam struct {
	BaseParam
	Params AddSecurityGroupRuleDetailParam `json:"params"` // 详细参数
}

