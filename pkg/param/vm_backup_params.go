// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmBackupDetailParam DeleteVmBackup详细参数
type DeleteVmBackupDetailParam struct {
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest []string `json:"backupStorageUuids,omitempty"`
	rest bool `json:"handleDependency,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVmBackupParam DeleteVmBackup请求参数
type DeleteVmBackupParam struct {
	BaseParam
	Params DeleteVmBackupDetailParam `json:"params"` // 详细参数
}

