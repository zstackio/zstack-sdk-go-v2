// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallRuleTemplateDetailParam DeleteFirewallRuleTemplate详细参数
type DeleteFirewallRuleTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleTemplateParam DeleteFirewallRuleTemplate请求参数
type DeleteFirewallRuleTemplateParam struct {
	BaseParam
	Params DeleteFirewallRuleTemplateDetailParam `json:"params"` // 详细参数
}

