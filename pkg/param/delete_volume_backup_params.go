// Copyright (c) ZStack.io, Inc.

package param

// DeleteVolumeBackupDetailParam DeleteVolumeBackup detail param
type DeleteVolumeBackupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	HandleDependency bool `json:"handleDependency,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVolumeBackupParam DeleteVolumeBackup request param
type DeleteVolumeBackupParam struct {
	BaseParam
	Params DeleteVolumeBackupDetailParam `json:"params"`
}
