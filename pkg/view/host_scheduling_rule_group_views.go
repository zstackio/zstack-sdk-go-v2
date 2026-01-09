// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostSchedulingRuleGroupInventoryView HostSchedulingRuleGroup
type HostSchedulingRuleGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateHostSchedulingRuleGroupEventView UpdateHostSchedulingRuleGroupEvent
type UpdateHostSchedulingRuleGroupEventView struct {
	Inventory HostSchedulingRuleGroupInventoryView `json:"inventory,omitempty"`
}

// QueryHostSchedulingRuleGroupView QueryHostSchedulingRuleGroup
type QueryHostSchedulingRuleGroupView struct {
	Inventories []HostSchedulingRuleGroupInventoryView `json:"inventories,omitempty"`
}

// CreateHostSchedulingRuleGroupEventView CreateHostSchedulingRuleGroupEvent
type CreateHostSchedulingRuleGroupEventView struct {
	Inventory HostSchedulingRuleGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteHostSchedulingRuleGroupEventView DeleteHostSchedulingRuleGroupEvent
type DeleteHostSchedulingRuleGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

