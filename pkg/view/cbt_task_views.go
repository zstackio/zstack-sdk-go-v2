// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CbtTaskInventoryView CbtTask
type CbtTaskInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []CbtTaskResourceRefInventoryView `json:"resourceRefs,omitempty"`
}

