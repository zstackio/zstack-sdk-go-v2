// Copyright (c) ZStack.io, Inc.

package param

// CreateRootVolumeTemplateFromVolumeBackupDetailParam CreateRootVolumeTemplateFromVolumeBackup detail param
type CreateRootVolumeTemplateFromVolumeBackupDetailParam struct {
	BackupUuid string `json:"backupUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeBackupParam CreateRootVolumeTemplateFromVolumeBackup request param
type CreateRootVolumeTemplateFromVolumeBackupParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromVolumeBackupDetailParam `json:"params"`
}
