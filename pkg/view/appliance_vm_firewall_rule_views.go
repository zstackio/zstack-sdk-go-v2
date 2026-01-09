// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ApplianceVmFirewallRuleInventoryView ApplianceVmFirewallRule
type ApplianceVmFirewallRuleInventoryView struct {
	ApplianceVmUuid *string `json:"applianceVmUuid,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	StartPort *int `json:"startPort,omitempty"`
	EndPort *int `json:"endPort,omitempty"`
	AllowCidr *string `json:"allowCidr,omitempty"`
	SourceIp *string `json:"sourceIp,omitempty"`
	DestIp *string `json:"destIp,omitempty"`
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

