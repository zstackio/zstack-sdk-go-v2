// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VolumeBackupInventoryView VolumeBackup
type VolumeBackupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	Mode string `json:"mode,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	BackupStorageRefs []VolumeBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

