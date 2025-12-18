// Copyright (c) ZStack.io, Inc.

package param

// SyncVolumeBackupDetailParam SyncVolumeBackup详细参数
type SyncVolumeBackupDetailParam struct {
	rest string `json:"imageStoreUuid" validate:"required"` // 必填
}

// SyncVolumeBackupParam SyncVolumeBackup请求参数
type SyncVolumeBackupParam struct {
	BaseParam
	Params SyncVolumeBackupDetailParam `json:"params"` // 详细参数
}

