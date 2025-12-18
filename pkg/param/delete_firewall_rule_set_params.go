// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallRuleSetDetailParam DeleteFirewallRuleSet detail param
type DeleteFirewallRuleSetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleSetParam DeleteFirewallRuleSet request param
type DeleteFirewallRuleSetParam struct {
	BaseParam
	Params DeleteFirewallRuleSetDetailParam `json:"params"`
}
