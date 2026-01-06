// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageBackupStorageRefInventoryView ImageBackupStorageRef
type ImageBackupStorageRefInventoryView struct {
	ImageUuid string `json:"imageUuid,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	Status string `json:"status,omitempty"`
	ExportMd5Sum string `json:"exportMd5Sum,omitempty"`
	ExportUrl string `json:"exportUrl,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

