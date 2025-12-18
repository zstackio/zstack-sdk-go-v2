// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcFirewallRuleSetInventoryView VpcFirewallRuleSet
type VpcFirewallRuleSetInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"actionType,omitempty"`
	rest string `json:"description,omitempty"`
	rest bool `json:"isDefault,omitempty"`
	rest bool `json:"isApplied,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []VpcFirewallRuleInventoryView `json:"rules,omitempty"`
}

