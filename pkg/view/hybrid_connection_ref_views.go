// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HybridConnectionRefInventoryView HybridConnectionRef
type HybridConnectionRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"connectionType,omitempty"`
	rest string `json:"connectionUuid,omitempty"`
	rest string `json:"direction,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

