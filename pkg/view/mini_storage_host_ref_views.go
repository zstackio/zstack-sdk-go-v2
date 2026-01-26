// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MiniStorageHostRefInventoryView MiniStorageHostRef
type MiniStorageHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	Status string `json:"status,omitempty"`
}

// QueryMiniStorageHostRefView QueryMiniStorageHostRef
type QueryMiniStorageHostRefView struct {
	Inventories []MiniStorageHostRefInventoryView `json:"inventories,omitempty"`
}

