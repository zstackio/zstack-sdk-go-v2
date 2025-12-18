// Copyright (c) ZStack.io, Inc.

package param

// ChangeFirewallRuleStateDetailParam ChangeFirewallRuleState detail param
type ChangeFirewallRuleStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	State string `json:"state" validate:"required"`
}

// ChangeFirewallRuleStateParam ChangeFirewallRuleState request param
type ChangeFirewallRuleStateParam struct {
	BaseParam
	Params ChangeFirewallRuleStateDetailParam `json:"params"`
}
