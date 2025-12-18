// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AccessControlListInventoryView AccessControlList
type AccessControlListInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Entries []AccessControlListEntryInventoryView `json:"entries,omitempty"`
}

