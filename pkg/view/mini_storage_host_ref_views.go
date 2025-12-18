// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MiniStorageHostRefInventoryView MiniStorageHostRef
type MiniStorageHostRefInventoryView struct {
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest int64 `json:"totalPhysicalCapacity,omitempty"`
	rest int64 `json:"availablePhysicalCapacity,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

