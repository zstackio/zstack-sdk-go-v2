// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SnapshotUsageInventoryView SnapshotUsage
type SnapshotUsageInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"SnapshotUuid,omitempty"`
	rest string `json:"SnapshotStatus,omitempty"`
	rest string `json:"SnapshotName,omitempty"`
	rest int64 `json:"SnapshotSize,omitempty"`
	rest string `json:"inventory,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

