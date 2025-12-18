// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LdapServerInventoryView LdapServer
type LdapServerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"base,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"scope,omitempty"`
	rest string `json:"encryption,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

