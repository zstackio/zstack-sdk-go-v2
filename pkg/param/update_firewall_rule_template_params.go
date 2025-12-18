// Copyright (c) ZStack.io, Inc.

package param

// UpdateFirewallRuleTemplateDetailParam UpdateFirewallRuleTemplate detail param
type UpdateFirewallRuleTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name" validate:"required"`
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
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateFirewallRuleTemplateParam UpdateFirewallRuleTemplate request param
type UpdateFirewallRuleTemplateParam struct {
	BaseParam
	Params UpdateFirewallRuleTemplateDetailParam `json:"params"`
}
