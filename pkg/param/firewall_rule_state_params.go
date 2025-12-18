// Copyright (c) ZStack.io, Inc.

package param

// ChangeFirewallRuleStateDetailParam ChangeFirewallRuleState详细参数
type ChangeFirewallRuleStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"state" validate:"required"` // 必填
}

// ChangeFirewallRuleStateParam ChangeFirewallRuleState请求参数
type ChangeFirewallRuleStateParam struct {
	BaseParam
	Params ChangeFirewallRuleStateDetailParam `json:"params"` // 详细参数
}

