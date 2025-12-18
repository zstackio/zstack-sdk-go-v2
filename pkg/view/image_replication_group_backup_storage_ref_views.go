// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageReplicationGroupBackupStorageRefInventoryView ImageReplicationGroupBackupStorageRef
type ImageReplicationGroupBackupStorageRefInventoryView struct {
	ReplicationGroupUuid string `json:"replicationGroupUuid,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

