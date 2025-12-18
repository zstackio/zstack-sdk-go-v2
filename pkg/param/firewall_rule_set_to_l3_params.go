// Copyright (c) ZStack.io, Inc.

package param

// AttachFirewallRuleSetToL3DetailParam AttachFirewallRuleSetToL3详细参数
type AttachFirewallRuleSetToL3DetailParam struct {
	rest string `json:"vpcFirewallUuid" validate:"required"` // 必填
	rest string `json:"l3Uuid" validate:"required"` // 必填
	rest string `json:"forward" validate:"required"` // 必填
	rest string `json:"ruleSetUuid" validate:"required"` // 必填
}

// AttachFirewallRuleSetToL3Param AttachFirewallRuleSetToL3请求参数
type AttachFirewallRuleSetToL3Param struct {
	BaseParam
	Params AttachFirewallRuleSetToL3DetailParam `json:"params"` // 详细参数
}

