// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AccountInventoryView Account
type AccountInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// UpdateAccountEventView UpdateAccountEvent
type UpdateAccountEventView struct {
	Inventory AccountInventoryView `json:"inventory,omitempty"`
}

// CleanInvalidLdapBindingEventView CleanInvalidLdapBindingEvent
type CleanInvalidLdapBindingEventView struct {
	Inventories []AccountInventoryView `json:"inventories,omitempty"`
}

// DeleteAccountEventView DeleteAccountEvent
type DeleteAccountEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateAccountEventView CreateAccountEvent
type CreateAccountEventView struct {
	Inventory AccountInventoryView `json:"inventory,omitempty"`
}

// QueryAccountView QueryAccount
type QueryAccountView struct {
	Inventories []AccountInventoryView `json:"inventories,omitempty"`
}

