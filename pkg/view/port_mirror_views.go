// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PortMirrorInventoryView PortMirror
type PortMirrorInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"mirrorNetworkUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []PortMirrorSessionInventoryView `json:"sessions,omitempty"`
}

