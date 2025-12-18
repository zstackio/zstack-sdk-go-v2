// Copyright (c) ZStack.io, Inc.

package param

// UpgradeBackupStorageCdpTasksDetailParam UpgradeBackupStorageCdpTasks详细参数
type UpgradeBackupStorageCdpTasksDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
}

// UpgradeBackupStorageCdpTasksParam UpgradeBackupStorageCdpTasks请求参数
type UpgradeBackupStorageCdpTasksParam struct {
	BaseParam
	Params UpgradeBackupStorageCdpTasksDetailParam `json:"params"` // 详细参数
}

