// Copyright (c) ZStack.io, Inc.

package param

// RevertVmFromVmBackupDetailParam RevertVmFromVmBackup详细参数
type RevertVmFromVmBackupDetailParam struct {
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid,omitempty"`
}

// RevertVmFromVmBackupParam RevertVmFromVmBackup请求参数
type RevertVmFromVmBackupParam struct {
	BaseParam
	Params RevertVmFromVmBackupDetailParam `json:"params"` // 详细参数
}

