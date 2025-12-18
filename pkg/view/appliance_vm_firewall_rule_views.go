// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ApplianceVmFirewallRuleInventoryView ApplianceVmFirewallRule
type ApplianceVmFirewallRuleInventoryView struct {
	rest string `json:"applianceVmUuid,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest int `json:"startPort,omitempty"`
	rest int `json:"endPort,omitempty"`
	rest string `json:"allowCidr,omitempty"`
	rest string `json:"sourceIp,omitempty"`
	rest string `json:"destIp,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

