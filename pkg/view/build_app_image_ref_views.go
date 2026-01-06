// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BuildAppImageRefInventoryView BuildAppImageRef
type BuildAppImageRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	BuildAppUuid string `json:"buildAppUuid,omitempty"`
	ImageName string `json:"imageName,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

