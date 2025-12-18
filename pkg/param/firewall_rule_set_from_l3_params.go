// Copyright (c) ZStack.io, Inc.

package param

// DetachFirewallRuleSetFromL3DetailParam DetachFirewallRuleSetFromL3详细参数
type DetachFirewallRuleSetFromL3DetailParam struct {
	rest string `json:"vpcFirewallUuid" validate:"required"` // 必填
	rest string `json:"l3Uuid" validate:"required"` // 必填
	rest string `json:"forward" validate:"required"` // 必填
	rest string `json:"ruleSetUuid" validate:"required"` // 必填
}

// DetachFirewallRuleSetFromL3Param DetachFirewallRuleSetFromL3请求参数
type DetachFirewallRuleSetFromL3Param struct {
	BaseParam
	Params DetachFirewallRuleSetFromL3DetailParam `json:"params"` // 详细参数
}

