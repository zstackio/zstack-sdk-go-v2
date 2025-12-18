// Copyright (c) ZStack.io, Inc.

package param

// CreateRootVolumeTemplateFromVolumeBackupDetailParam CreateRootVolumeTemplateFromVolumeBackup详细参数
type CreateRootVolumeTemplateFromVolumeBackupDetailParam struct {
	rest string `json:"backupUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"guestOsType,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest bool `json:"system,omitempty"`
	rest bool `json:"virtio,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeBackupParam CreateRootVolumeTemplateFromVolumeBackup请求参数
type CreateRootVolumeTemplateFromVolumeBackupParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromVolumeBackupDetailParam `json:"params"` // 详细参数
}

