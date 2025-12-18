// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmBackupDetailParam DeleteVmBackup detail param
type DeleteVmBackupDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	HandleDependency bool `json:"handleDependency,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmBackupParam DeleteVmBackup request param
type DeleteVmBackupParam struct {
	BaseParam
	Params DeleteVmBackupDetailParam `json:"params"`
}
