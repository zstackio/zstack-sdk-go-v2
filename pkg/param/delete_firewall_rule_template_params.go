// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallRuleTemplateDetailParam DeleteFirewallRuleTemplate detail param
type DeleteFirewallRuleTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleTemplateParam DeleteFirewallRuleTemplate request param
type DeleteFirewallRuleTemplateParam struct {
	BaseParam
	Params DeleteFirewallRuleTemplateDetailParam `json:"params"`
}
