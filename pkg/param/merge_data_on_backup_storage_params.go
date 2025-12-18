// Copyright (c) ZStack.io, Inc.

package param

// MergeDataOnBackupStorageDetailParam MergeDataOnBackupStorage detail param
type MergeDataOnBackupStorageDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
}

// MergeDataOnBackupStorageParam MergeDataOnBackupStorage request param
type MergeDataOnBackupStorageParam struct {
	BaseParam
	Params MergeDataOnBackupStorageDetailParam `json:"params"`
}
