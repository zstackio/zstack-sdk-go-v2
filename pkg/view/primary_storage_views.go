// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PrimaryStorageInventoryView PrimaryStorage
type PrimaryStorageInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	TotalCapacity *int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity *int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity *int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity *int64 `json:"availablePhysicalCapacity,omitempty"`
	SystemUsedCapacity *int64 `json:"systemUsedCapacity,omitempty"`
	Type *string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	MountPath *string `json:"mountPath,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// AttachPrimaryStorageToClusterEventView AttachPrimaryStorageToClusterEvent
type AttachPrimaryStorageToClusterEventView struct {
	Inventory PrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// GetPrimaryStorageCandidatesForVmMigrationView GetPrimaryStorageCandidatesForVmMigration
type GetPrimaryStorageCandidatesForVmMigrationView struct {
	Inventories []PrimaryStorageInventoryView `json:"inventories,omitempty"`
}

// GetPrimaryStorageCandidatesForVolumeMigrationView GetPrimaryStorageCandidatesForVolumeMigration
type GetPrimaryStorageCandidatesForVolumeMigrationView struct {
	Inventories []PrimaryStorageInventoryView `json:"inventories,omitempty"`
}

// ChangePrimaryStorageStateEventView ChangePrimaryStorageStateEvent
type ChangePrimaryStorageStateEventView struct {
	Inventory PrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// ReconnectPrimaryStorageEventView ReconnectPrimaryStorageEvent
type ReconnectPrimaryStorageEventView struct {
	Inventory PrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// DetachPrimaryStorageFromClusterEventView DetachPrimaryStorageFromClusterEvent
type DetachPrimaryStorageFromClusterEventView struct {
	Inventory PrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// DeletePrimaryStorageEventView DeletePrimaryStorageEvent
type DeletePrimaryStorageEventView struct {
	Success bool `json:"success,omitempty"`
}

