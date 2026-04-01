// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VolumeBackupInventoryView VolumeBackup
type VolumeBackupInventoryView struct {
	BaseInfoView
	BaseTimeView
	VolumeUuid string `json:"volumeUuid,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	Mode string `json:"mode,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	BackupStorageRefs []VolumeBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

// CreateVmBackupEventView CreateVmBackupEvent
type CreateVmBackupEventView struct {
	Inventories []VolumeBackupInventoryView `json:"inventories,omitempty"`
}

// RecoverVmBackupFromImageStoreBackupStorageEventView RecoverVmBackupFromImageStoreBackupStorageEvent
type RecoverVmBackupFromImageStoreBackupStorageEventView struct {
	Inventories []VolumeBackupInventoryView `json:"inventories,omitempty"`
}

// SyncVmBackupFromImageStoreBackupStorageEventView SyncVmBackupFromImageStoreBackupStorageEvent
type SyncVmBackupFromImageStoreBackupStorageEventView struct {
	Inventories []VolumeBackupInventoryView `json:"inventories,omitempty"`
}

// CreateVolumeBackupEventView CreateVolumeBackupEvent
type CreateVolumeBackupEventView struct {
	Inventory VolumeBackupInventoryView `json:"inventory,omitempty"`
}

// SyncVolumeBackupEventView SyncVolumeBackupEvent
type SyncVolumeBackupEventView struct {
	Result SyncBackupResultView `json:"result,omitempty"`
}

// RecoverBackupFromImageStoreBackupStorageEventView RecoverBackupFromImageStoreBackupStorageEvent
type RecoverBackupFromImageStoreBackupStorageEventView struct {
	Inventory VolumeBackupInventoryView `json:"inventory,omitempty"`
}

// QueryVolumeBackupView QueryVolumeBackup
type QueryVolumeBackupView struct {
	Inventories []VolumeBackupInventoryView `json:"inventories,omitempty"`
}

// DeleteVolumeBackupEventView DeleteVolumeBackupEvent
type DeleteVolumeBackupEventView struct {
	Success bool `json:"success,omitempty"`
}

// SyncBackupFromImageStoreBackupStorageEventView SyncBackupFromImageStoreBackupStorageEvent
type SyncBackupFromImageStoreBackupStorageEventView struct {
	Inventory VolumeBackupInventoryView `json:"inventory,omitempty"`
}

