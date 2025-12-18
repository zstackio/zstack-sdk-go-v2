// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AccessControlListInventoryView AccessControlList
type AccessControlListInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []AccessControlListEntryInventoryView `json:"entries,omitempty"`
}

