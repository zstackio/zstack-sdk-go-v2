// Copyright (c) ZStack.io, Inc.

package param

// SyncVmBackupDetailParam SyncVmBackup详细参数
type SyncVmBackupDetailParam struct {
	rest string `json:"imageStoreUuid" validate:"required"` // 必填
}

// SyncVmBackupParam SyncVmBackup请求参数
type SyncVmBackupParam struct {
	BaseParam
	Params SyncVmBackupDetailParam `json:"params"` // 详细参数
}

