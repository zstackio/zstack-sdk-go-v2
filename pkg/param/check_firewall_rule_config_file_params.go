// Copyright (c) ZStack.io, Inc.

package param

// CheckFirewallRuleConfigFileDetailParam CheckFirewallRuleConfigFile详细参数
type CheckFirewallRuleConfigFileDetailParam struct {
	rest string `json:"ruleInfo" validate:"required"` // 必填
}

// CheckFirewallRuleConfigFileParam CheckFirewallRuleConfigFile请求参数
type CheckFirewallRuleConfigFileParam struct {
	BaseParam
	Params CheckFirewallRuleConfigFileDetailParam `json:"params"` // 详细参数
}

