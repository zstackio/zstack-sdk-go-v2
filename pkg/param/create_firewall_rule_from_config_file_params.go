// Copyright (c) ZStack.io, Inc.

package param

// CreateFirewallRuleFromConfigFileDetailParam CreateFirewallRuleFromConfigFile detail param
type CreateFirewallRuleFromConfigFileDetailParam struct {
	RuleInfo string `json:"ruleInfo" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleFromConfigFileParam CreateFirewallRuleFromConfigFile request param
type CreateFirewallRuleFromConfigFileParam struct {
	BaseParam
	Params CreateFirewallRuleFromConfigFileDetailParam `json:"params"`
}
