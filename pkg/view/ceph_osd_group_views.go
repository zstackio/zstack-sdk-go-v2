// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CephOsdGroupInventoryView CephOsdGroup
type CephOsdGroupInventoryView struct {
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"osds,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest int64 `json:"availablePhysicalCapacity,omitempty"`
	rest int64 `json:"totalPhysicalCapacity,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"uuid,omitempty"`
}

