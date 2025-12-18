// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LdapAccountRefInventoryView LdapAccountRef
type LdapAccountRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"ldapUid,omitempty"`
	rest string `json:"ldapServerUuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

