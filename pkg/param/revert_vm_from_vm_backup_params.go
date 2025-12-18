// Copyright (c) ZStack.io, Inc.

package param

// RevertVmFromVmBackupDetailParam RevertVmFromVmBackup detail param
type RevertVmFromVmBackupDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
}

// RevertVmFromVmBackupParam RevertVmFromVmBackup request param
type RevertVmFromVmBackupParam struct {
	BaseParam
	Params RevertVmFromVmBackupDetailParam `json:"params"`
}
