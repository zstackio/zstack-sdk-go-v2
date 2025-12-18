// Copyright (c) ZStack.io, Inc.

package param

// AddBackupStoragesToReplicationGroupDetailParam AddBackupStoragesToReplicationGroup detail param
type AddBackupStoragesToReplicationGroupDetailParam struct {
	ReplicationGroupUuid string `json:"replicationGroupUuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBackupStoragesToReplicationGroupParam AddBackupStoragesToReplicationGroup request param
type AddBackupStoragesToReplicationGroupParam struct {
	BaseParam
	Params AddBackupStoragesToReplicationGroupDetailParam `json:"params"`
}
