// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MiniStorageInventoryView MiniStorage
type MiniStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	MiniStorageType string `json:"miniStorageType,omitempty"`
	DiskIdentifier string `json:"diskIdentifier,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	SystemUsedCapacity int64 `json:"systemUsedCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// QueryMiniStorageView QueryMiniStorage
type QueryMiniStorageView struct {
	Inventories []MiniStorageInventoryView `json:"inventories,omitempty"`
}

// AddPrimaryStorageEventView AddPrimaryStorageEvent
type AddPrimaryStorageEventView struct {
	Inventory PrimaryStorageInventoryView `json:"inventory,omitempty"`
}

