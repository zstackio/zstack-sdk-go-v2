// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LdapAccountRefInventoryView LdapAccountRef
type LdapAccountRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	LdapUid *string `json:"ldapUid,omitempty"`
	LdapServerUuid *string `json:"ldapServerUuid,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryLdapBindingView QueryLdapBinding
type QueryLdapBindingView struct {
	Inventories []LdapAccountRefInventoryView `json:"inventories,omitempty"`
}

// CreateLdapBindingEventView CreateLdapBindingEvent
type CreateLdapBindingEventView struct {
	Inventory LdapAccountRefInventoryView `json:"inventory,omitempty"`
}

