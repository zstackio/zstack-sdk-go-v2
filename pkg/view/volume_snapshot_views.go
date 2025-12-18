// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VolumeSnapshotInventoryView VolumeSnapshot
type VolumeSnapshotInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	TreeUuid string `json:"treeUuid,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	PrimaryStorageInstallPath string `json:"primaryStorageInstallPath,omitempty"`
	VolumeType string `json:"volumeType,omitempty"`
	Format string `json:"format,omitempty"`
	Latest bool `json:"latest,omitempty"`
	Size int64 `json:"size,omitempty"`
	Distance int `json:"distance,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	BackupStorageRefs []VolumeSnapshotBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
}

