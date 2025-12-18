// Copyright (c) ZStack.io, Inc.

package param

// UpgradeBackupStorageCdpTasksDetailParam UpgradeBackupStorageCdpTasks detail param
type UpgradeBackupStorageCdpTasksDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
}

// UpgradeBackupStorageCdpTasksParam UpgradeBackupStorageCdpTasks request param
type UpgradeBackupStorageCdpTasksParam struct {
	BaseParam
	Params UpgradeBackupStorageCdpTasksDetailParam `json:"params"`
}
