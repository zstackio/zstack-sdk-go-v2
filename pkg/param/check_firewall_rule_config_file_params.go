// Copyright (c) ZStack.io, Inc.

package param

// CheckFirewallRuleConfigFileDetailParam CheckFirewallRuleConfigFile detail param
type CheckFirewallRuleConfigFileDetailParam struct {
	RuleInfo string `json:"ruleInfo" validate:"required"`
}

// CheckFirewallRuleConfigFileParam CheckFirewallRuleConfigFile request param
type CheckFirewallRuleConfigFileParam struct {
	BaseParam
	Params CheckFirewallRuleConfigFileDetailParam `json:"params"`
}
