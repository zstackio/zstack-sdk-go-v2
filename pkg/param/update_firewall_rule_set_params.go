// Copyright (c) ZStack.io, Inc.

package param

// UpdateFirewallRuleSetDetailParam UpdateFirewallRuleSet detail param
type UpdateFirewallRuleSetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ActionType string `json:"actionType,omitempty"`
}

// UpdateFirewallRuleSetParam UpdateFirewallRuleSet request param
type UpdateFirewallRuleSetParam struct {
	BaseParam
	Params UpdateFirewallRuleSetDetailParam `json:"params"`
}
