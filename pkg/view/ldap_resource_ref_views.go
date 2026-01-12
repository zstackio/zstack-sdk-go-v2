// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LdapResourceRefInventoryView LdapResourceRef
type LdapResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	LdapUid *string `json:"ldapUid,omitempty"`
	LdapServerUuid *string `json:"ldapServerUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

// QueryIAM2LdapBindingView QueryIAM2LdapBinding
type QueryIAM2LdapBindingView struct {
	Inventories []LdapResourceRefInventoryView `json:"inventories,omitempty"`
}

// CreateIAM2VirtualIDLdapBindingEventView CreateIAM2VirtualIDLdapBindingEvent
type CreateIAM2VirtualIDLdapBindingEventView struct {
	Inventory LdapResourceRefInventoryView `json:"inventory,omitempty"`
}

// CreateIAM2VirtualIDFromLdapUidEventView CreateIAM2VirtualIDFromLdapUidEvent
type CreateIAM2VirtualIDFromLdapUidEventView struct {
	Inventory LdapResourceRefInventoryView `json:"inventory,omitempty"`
}

