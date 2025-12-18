// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ZoneInventoryView Zone
type ZoneInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"isDefault,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

