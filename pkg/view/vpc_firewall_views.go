// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcFirewallInventoryView VpcFirewall
type VpcFirewallInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest []VpcFirewallRuleSetL3RefInventoryView `json:"refs,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"description,omitempty"`
}

