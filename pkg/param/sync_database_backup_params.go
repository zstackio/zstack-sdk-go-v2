// Copyright (c) ZStack.io, Inc.

package param

// SyncDatabaseBackupDetailParam SyncDatabaseBackup detail param
type SyncDatabaseBackupDetailParam struct {
	ImageStoreUuid string `json:"imageStoreUuid" validate:"required"`
}

// SyncDatabaseBackupParam SyncDatabaseBackup request param
type SyncDatabaseBackupParam struct {
	BaseParam
	Params SyncDatabaseBackupDetailParam `json:"params"`
}
