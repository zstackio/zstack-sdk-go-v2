// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LdapResourceRefInventoryView LdapResourceRef
type LdapResourceRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"ldapUid,omitempty"`
	rest string `json:"ldapServerUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

