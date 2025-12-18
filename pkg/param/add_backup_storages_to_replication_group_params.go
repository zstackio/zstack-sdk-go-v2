// Copyright (c) ZStack.io, Inc.

package param

// AddBackupStoragesToReplicationGroupDetailParam AddBackupStoragesToReplicationGroup详细参数
type AddBackupStoragesToReplicationGroupDetailParam struct {
	rest string `json:"replicationGroupUuid" validate:"required"` // 必填
	rest []string `json:"backupStorageUuids" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddBackupStoragesToReplicationGroupParam AddBackupStoragesToReplicationGroup请求参数
type AddBackupStoragesToReplicationGroupParam struct {
	BaseParam
	Params AddBackupStoragesToReplicationGroupDetailParam `json:"params"` // 详细参数
}

