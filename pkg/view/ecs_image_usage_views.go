// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EcsImageUsageInventoryView EcsImageUsage
type EcsImageUsageInventoryView struct {
	rest int `json:"id,omitempty"`
	rest string `json:"ecsImageUuid,omitempty"`
	rest string `json:"snapshotUuidOfCreatedImage,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

