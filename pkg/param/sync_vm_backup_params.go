// Copyright (c) ZStack.io, Inc.

package param

// SyncVmBackupDetailParam SyncVmBackup detail param
type SyncVmBackupDetailParam struct {
	ImageStoreUuid string `json:"imageStoreUuid" validate:"required"`
}

// SyncVmBackupParam SyncVmBackup request param
type SyncVmBackupParam struct {
	BaseParam
	Params SyncVmBackupDetailParam `json:"params"`
}
