// Copyright (c) ZStack.io, Inc.

package view

// LogInByLdapView LogInByLdap
type LogInByLdapView struct {
	Inventory SessionInventoryView `json:"inventory,omitempty"`
	AccountInventory AccountInventoryView `json:"accountInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

