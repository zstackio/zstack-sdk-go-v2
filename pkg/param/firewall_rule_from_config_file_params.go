// Copyright (c) ZStack.io, Inc.

package param

// CreateFirewallRuleFromConfigFileDetailParam CreateFirewallRuleFromConfigFile详细参数
type CreateFirewallRuleFromConfigFileDetailParam struct {
	rest string `json:"ruleInfo" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleFromConfigFileParam CreateFirewallRuleFromConfigFile请求参数
type CreateFirewallRuleFromConfigFileParam struct {
	BaseParam
	Params CreateFirewallRuleFromConfigFileDetailParam `json:"params"` // 详细参数
}

