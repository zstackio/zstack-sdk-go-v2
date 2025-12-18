// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ImageBackupStorageRefInventoryView ImageBackupStorageRef
type ImageBackupStorageRefInventoryView struct {
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"exportMd5Sum,omitempty"`
	rest string `json:"exportUrl,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

