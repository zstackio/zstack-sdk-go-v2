// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeFromVolumeBackupDetailParam CreateDataVolumeFromVolumeBackup详细参数
type CreateDataVolumeFromVolumeBackupDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"backupUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeBackupParam CreateDataVolumeFromVolumeBackup请求参数
type CreateDataVolumeFromVolumeBackupParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeBackupDetailParam `json:"params"` // 详细参数
}

