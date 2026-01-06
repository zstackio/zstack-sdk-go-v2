// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BackupStorageZoneRefInventoryView BackupStorageZoneRef
type BackupStorageZoneRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

