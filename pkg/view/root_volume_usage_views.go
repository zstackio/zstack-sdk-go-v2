// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RootVolumeUsageInventoryView RootVolumeUsage
type RootVolumeUsageInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest int64 `json:"dateInLong,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"volumeStatus,omitempty"`
	rest int64 `json:"volumeName,omitempty"`
	rest int64 `json:"volumeSize,omitempty"`
	rest string `json:"inventory,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

