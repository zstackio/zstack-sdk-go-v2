// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallRuleSetDetailParam DeleteFirewallRuleSet详细参数
type DeleteFirewallRuleSetDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleSetParam DeleteFirewallRuleSet请求参数
type DeleteFirewallRuleSetParam struct {
	BaseParam
	Params DeleteFirewallRuleSetDetailParam `json:"params"` // 详细参数
}

