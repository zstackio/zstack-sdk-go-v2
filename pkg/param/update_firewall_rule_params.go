// Copyright (c) ZStack.io, Inc.

package param

// UpdateFirewallRuleDetailParam UpdateFirewallRule detail param
type UpdateFirewallRuleDetailParam struct {
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	Action string `json:"action" validate:"required"`
	Protocol string `json:"protocol,omitempty"`
	DestPort string `json:"destPort,omitempty"`
	SourcePort string `json:"sourcePort,omitempty"`
	SourceIp string `json:"sourceIp,omitempty"`
	DestIp string `json:"destIp,omitempty"`
	AllowStates string `json:"allowStates,omitempty"`
	TcpFlag string `json:"tcpFlag,omitempty"`
	IcmpTypeName string `json:"icmpTypeName,omitempty"`
	RuleNumber int `json:"ruleNumber" validate:"required"`
	EnableLog bool `json:"enableLog,omitempty"`
	State string `json:"state" validate:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateFirewallRuleParam UpdateFirewallRule request param
type UpdateFirewallRuleParam struct {
	BaseParam
	Params UpdateFirewallRuleDetailParam `json:"params"`
}
