// Copyright (c) ZStack.io, Inc.

package param

// RevertVolumeFromVolumeBackupDetailParam RevertVolumeFromVolumeBackup detail param
type RevertVolumeFromVolumeBackupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
}

// RevertVolumeFromVolumeBackupParam RevertVolumeFromVolumeBackup request param
type RevertVolumeFromVolumeBackupParam struct {
	BaseParam
	Params RevertVolumeFromVolumeBackupDetailParam `json:"params"`
}
