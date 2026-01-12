// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ExternalPrimaryStorageSpaceInventoryView ExternalPrimaryStorageSpace
type ExternalPrimaryStorageSpaceInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	LocationUrl *string `json:"locationUrl,omitempty"`
	Type *string `json:"type,omitempty"`
	AvailableCapacity *int64 `json:"availableCapacity,omitempty"`
	TotalCapacity *int64 `json:"totalCapacity,omitempty"`
	AvailablePhysicalCapacity *int64 `json:"availablePhysicalCapacity,omitempty"`
	TotalPhysicalCapacity *int64 `json:"totalPhysicalCapacity,omitempty"`
}

