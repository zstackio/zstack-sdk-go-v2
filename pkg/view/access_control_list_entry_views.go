// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AccessControlListEntryInventoryView AccessControlListEntry
type AccessControlListEntryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"aclUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"domain,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"ipEntries,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

