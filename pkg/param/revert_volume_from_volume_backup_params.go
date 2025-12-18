// Copyright (c) ZStack.io, Inc.

package param

// RevertVolumeFromVolumeBackupDetailParam RevertVolumeFromVolumeBackup详细参数
type RevertVolumeFromVolumeBackupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid,omitempty"`
}

// RevertVolumeFromVolumeBackupParam RevertVolumeFromVolumeBackup请求参数
type RevertVolumeFromVolumeBackupParam struct {
	BaseParam
	Params RevertVolumeFromVolumeBackupDetailParam `json:"params"` // 详细参数
}

