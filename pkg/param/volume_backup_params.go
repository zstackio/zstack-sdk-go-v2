// Copyright (c) ZStack.io, Inc.

package param

// CreateVolumeBackupDetailParam CreateVolumeBackup详细参数
type CreateVolumeBackupDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"mode,omitempty"`
	rest int64 `json:"volumeReadBandwidth,omitempty"`
	rest int64 `json:"volumeWriteBandwidth,omitempty"`
	rest int64 `json:"networkReadBandwidth,omitempty"`
	rest int64 `json:"networkWriteBandwidth,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVolumeBackupParam CreateVolumeBackup请求参数
type CreateVolumeBackupParam struct {
	BaseParam
	Params CreateVolumeBackupDetailParam `json:"params"` // 详细参数
}

