// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BackupStorageZoneRefInventoryView BackupStorageZoneRef
type BackupStorageZoneRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

