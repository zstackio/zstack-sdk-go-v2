// Copyright (c) ZStack.io, Inc.

package param

// AttachFirewallRuleSetToL3DetailParam AttachFirewallRuleSetToL3 detail param
type AttachFirewallRuleSetToL3DetailParam struct {
	VpcFirewallUuid string `json:"vpcFirewallUuid" validate:"required"`
	L3Uuid string `json:"l3Uuid" validate:"required"`
	Forward string `json:"forward" validate:"required"`
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
}

// AttachFirewallRuleSetToL3Param AttachFirewallRuleSetToL3 request param
type AttachFirewallRuleSetToL3Param struct {
	BaseParam
	Params AttachFirewallRuleSetToL3DetailParam `json:"params"`
}
