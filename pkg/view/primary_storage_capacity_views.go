// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PrimaryStorageCapacityInventoryView PrimaryStorageCapacity
type PrimaryStorageCapacityInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest int64 `json:"totalPhysicalCapacity,omitempty"`
	rest int64 `json:"availablePhysicalCapacity,omitempty"`
	rest int64 `json:"systemUsedCapacity,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

