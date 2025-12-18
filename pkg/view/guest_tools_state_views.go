// Copyright (c) ZStack.io, Inc.

package view

import "time"

// GuestToolsStateInventoryView GuestToolsState
type GuestToolsStateInventoryView struct {
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"qgaState,omitempty"`
	rest string `json:"zwatchState,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"osType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

