// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LdapServerInventoryView LdapServer
type LdapServerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	Base *string `json:"base,omitempty"`
	Username *string `json:"username,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Encryption *string `json:"encryption,omitempty"`
}

// AddLdapServerEventView AddLdapServerEvent
type AddLdapServerEventView struct {
	Inventory LdapServerInventoryView `json:"inventory,omitempty"`
}

// QueryLdapServerView QueryLdapServer
type QueryLdapServerView struct {
	Inventories []LdapServerInventoryView `json:"inventories,omitempty"`
}

// SyncLdapServerEventView SyncLdapServerEvent
type SyncLdapServerEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// DeleteLdapServerEventView DeleteLdapServerEvent
type DeleteLdapServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateLdapServerEventView UpdateLdapServerEvent
type UpdateLdapServerEventView struct {
	Inventory LdapServerInventoryView `json:"inventory,omitempty"`
}

