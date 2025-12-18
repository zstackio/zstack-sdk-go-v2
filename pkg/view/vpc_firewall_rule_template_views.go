// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcFirewallRuleTemplateInventoryView VpcFirewallRuleTemplate
type VpcFirewallRuleTemplateInventoryView struct {
	rest string `json:"action,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"destPort,omitempty"`
	rest string `json:"sourcePort,omitempty"`
	rest string `json:"sourceIp,omitempty"`
	rest string `json:"destIp,omitempty"`
	rest string `json:"allowStates,omitempty"`
	rest string `json:"tcpFlag,omitempty"`
	rest string `json:"icmpTypeName,omitempty"`
	rest int `json:"ruleNumber,omitempty"`
	rest bool `json:"enableLog,omitempty"`
	rest string `json:"state,omitempty"`
	rest bool `json:"isDefault,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
}

