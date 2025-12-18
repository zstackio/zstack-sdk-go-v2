// Copyright (c) ZStack.io, Inc.

package param

// SyncDatabaseBackupDetailParam SyncDatabaseBackup详细参数
type SyncDatabaseBackupDetailParam struct {
	rest string `json:"imageStoreUuid" validate:"required"` // 必填
}

// SyncDatabaseBackupParam SyncDatabaseBackup请求参数
type SyncDatabaseBackupParam struct {
	BaseParam
	Params SyncDatabaseBackupDetailParam `json:"params"` // 详细参数
}

