// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AccessControlRuleInventoryView AccessControlRule
type AccessControlRuleInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Rule *string `json:"rule,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// AddAccessControlRuleEventView AddAccessControlRuleEvent
type AddAccessControlRuleEventView struct {
	Inventory AccessControlRuleInventoryView `json:"inventory,omitempty"`
}

// UpdateAccessControlRuleEventView UpdateAccessControlRuleEvent
type UpdateAccessControlRuleEventView struct {
	Inventory AccessControlRuleInventoryView `json:"inventory,omitempty"`
}

// DeleteAccessControlRuleEventView DeleteAccessControlRuleEvent
type DeleteAccessControlRuleEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryAccessControlRuleView QueryAccessControlRule
type QueryAccessControlRuleView struct {
	Inventories []AccessControlRuleInventoryView `json:"inventories,omitempty"`
}

