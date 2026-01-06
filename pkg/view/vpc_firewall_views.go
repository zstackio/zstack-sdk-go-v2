// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcFirewallInventoryView VpcFirewall
type VpcFirewallInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Refs []VpcFirewallRuleSetL3RefInventoryView `json:"refs,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVpcFirewallEventView UpdateVpcFirewallEvent
type UpdateVpcFirewallEventView struct {
	Inventory VpcFirewallInventoryView `json:"inventory,omitempty"`
}

// CreateVpcFirewallEventView CreateVpcFirewallEvent
type CreateVpcFirewallEventView struct {
	Inventory VpcFirewallInventoryView `json:"inventory,omitempty"`
}

// RefreshFirewallEventView RefreshFirewallEvent
type RefreshFirewallEventView struct {
	Inventory VpcFirewallInventoryView `json:"inventory,omitempty"`
}

// QueryVpcFirewallView QueryVpcFirewall
type QueryVpcFirewallView struct {
	Inventories []VpcFirewallInventoryView `json:"inventories,omitempty"`
}

