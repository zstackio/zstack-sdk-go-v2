// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LdapServerInventoryView LdapServer
type LdapServerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	Base string `json:"base,omitempty"`
	Username string `json:"username,omitempty"`
	Scope string `json:"scope,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

