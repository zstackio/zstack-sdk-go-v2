// Copyright (c) ZStack.io, Inc.

package param

// DetachFirewallRuleSetFromL3DetailParam DetachFirewallRuleSetFromL3 detail param
type DetachFirewallRuleSetFromL3DetailParam struct {
	VpcFirewallUuid string `json:"vpcFirewallUuid" validate:"required"`
	L3Uuid string `json:"l3Uuid" validate:"required"`
	Forward string `json:"forward" validate:"required"`
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
}

// DetachFirewallRuleSetFromL3Param DetachFirewallRuleSetFromL3 request param
type DetachFirewallRuleSetFromL3Param struct {
	BaseParam
	Params DetachFirewallRuleSetFromL3DetailParam `json:"params"`
}
