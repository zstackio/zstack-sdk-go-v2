// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcFirewallRuleSetInventoryView VpcFirewallRuleSet
type VpcFirewallRuleSetInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	ActionType string `json:"actionType,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	IsApplied bool `json:"isApplied,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Rules []VpcFirewallRuleInventoryView `json:"rules,omitempty"`
}

