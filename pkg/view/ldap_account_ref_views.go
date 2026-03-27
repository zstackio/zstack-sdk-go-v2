// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LdapAccountRefInventoryView LdapAccountRef
type LdapAccountRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	LdapUid string `json:"ldapUid,omitempty"`
	LdapServerUuid string `json:"ldapServerUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// QueryLdapBindingView QueryLdapBinding
type QueryLdapBindingView struct {
	Inventories []LdapAccountRefInventoryView `json:"inventories,omitempty"`
}

// CreateLdapBindingEventView CreateLdapBindingEvent
type CreateLdapBindingEventView struct {
	Inventory LdapAccountRefInventoryView `json:"inventory,omitempty"`
}

