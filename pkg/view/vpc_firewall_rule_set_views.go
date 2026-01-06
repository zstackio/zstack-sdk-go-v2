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
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Rules []VpcFirewallRuleInventoryView `json:"rules,omitempty"`
}

// CreateFirewallRuleFromConfigFileEventView CreateFirewallRuleFromConfigFileEvent
type CreateFirewallRuleFromConfigFileEventView struct {
	Inventory []VpcFirewallRuleSetInventoryView `json:"inventory,omitempty"`
}

// QueryFirewallRuleSetView QueryFirewallRuleSet
type QueryFirewallRuleSetView struct {
	Inventories []VpcFirewallRuleSetInventoryView `json:"inventories,omitempty"`
}

// CreateFirewallRuleSetEventView CreateFirewallRuleSetEvent
type CreateFirewallRuleSetEventView struct {
	Inventory VpcFirewallRuleSetInventoryView `json:"inventory,omitempty"`
}

// ApplyRuleSetChangesEventView ApplyRuleSetChangesEvent
type ApplyRuleSetChangesEventView struct {
	Inventory VpcFirewallRuleSetInventoryView `json:"inventory,omitempty"`
}

// UpdateFirewallRuleSetEventView UpdateFirewallRuleSetEvent
type UpdateFirewallRuleSetEventView struct {
	Inventory VpcFirewallRuleSetInventoryView `json:"inventory,omitempty"`
}

