// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PortMirrorSessionInventoryView PortMirrorSession
type PortMirrorSessionInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest int64 `json:"internalId,omitempty"`
	rest string `json:"srcEndPoint,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"dstEndPoint,omitempty"`
	rest string `json:"portMirrorUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

