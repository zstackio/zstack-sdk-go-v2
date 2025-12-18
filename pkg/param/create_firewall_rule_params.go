// Copyright (c) ZStack.io, Inc.

package param

// CreateFirewallRuleDetailParam CreateFirewallRule detail param
type CreateFirewallRuleDetailParam struct {
	RuleSetUuid string `json:"ruleSetUuid" validate:"required"`
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
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleParam CreateFirewallRule request param
type CreateFirewallRuleParam struct {
	BaseParam
	Params CreateFirewallRuleDetailParam `json:"params"`
}
