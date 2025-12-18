// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallRuleDetailParam DeleteFirewallRule detail param
type DeleteFirewallRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFirewallRuleParam DeleteFirewallRule request param
type DeleteFirewallRuleParam struct {
	BaseParam
	Params DeleteFirewallRuleDetailParam `json:"params"`
}
