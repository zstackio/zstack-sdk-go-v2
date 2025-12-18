// Copyright (c) ZStack.io, Inc.

package param

// CreateFirewallRuleTemplateDetailParam CreateFirewallRuleTemplate detail param
type CreateFirewallRuleTemplateDetailParam struct {
	Action string `json:"action" validate:"required"`
	Protocol string `json:"protocol,omitempty"`
	Name string `json:"name" validate:"required"`
	DestPort string `json:"destPort,omitempty"`
	SourcePort string `json:"sourcePort,omitempty"`
	SourceIp string `json:"sourceIp,omitempty"`
	DestIp string `json:"destIp,omitempty"`
	AllowStates string `json:"allowStates,omitempty"`
	TcpFlag string `json:"tcpFlag,omitempty"`
	IcmpTypeName string `json:"icmpTypeName,omitempty"`
	RuleNumber int `json:"ruleNumber" validate:"required"`
	EnableLog bool `json:"enableLog,omitempty"`
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFirewallRuleTemplateParam CreateFirewallRuleTemplate request param
type CreateFirewallRuleTemplateParam struct {
	BaseParam
	Params CreateFirewallRuleTemplateDetailParam `json:"params"`
}
