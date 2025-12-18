// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ImageReplicationGroupBackupStorageRefInventoryView ImageReplicationGroupBackupStorageRef
type ImageReplicationGroupBackupStorageRefInventoryView struct {
	rest string `json:"replicationGroupUuid,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

