// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TagPatternInventoryView TagPattern
type TagPatternInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"value,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"color,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

