// Copyright (c) ZStack.io, Inc.

package param

// UpdateFirewallRuleDetailParam UpdateFirewallRule详细参数
type UpdateFirewallRuleDetailParam struct {
	rest string `json:"ruleSetUuid" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"action" validate:"required"` // 必填
	rest string `json:"protocol,omitempty"`
	rest string `json:"destPort,omitempty"`
	rest string `json:"sourcePort,omitempty"`
	rest string `json:"sourceIp,omitempty"`
	rest string `json:"destIp,omitempty"`
	rest string `json:"allowStates,omitempty"`
	rest string `json:"tcpFlag,omitempty"`
	rest string `json:"icmpTypeName,omitempty"`
	rest int `json:"ruleNumber" validate:"required"` // 必填
	rest bool `json:"enableLog,omitempty"`
	rest string `json:"state" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
}

// UpdateFirewallRuleParam UpdateFirewallRule请求参数
type UpdateFirewallRuleParam struct {
	BaseParam
	Params UpdateFirewallRuleDetailParam `json:"params"` // 详细参数
}

