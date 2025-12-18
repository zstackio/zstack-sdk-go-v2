// Copyright (c) ZStack.io, Inc.

package param

// SyncVolumeBackupDetailParam SyncVolumeBackup detail param
type SyncVolumeBackupDetailParam struct {
	ImageStoreUuid string `json:"imageStoreUuid" validate:"required"`
}

// SyncVolumeBackupParam SyncVolumeBackup request param
type SyncVolumeBackupParam struct {
	BaseParam
	Params SyncVolumeBackupDetailParam `json:"params"`
}
