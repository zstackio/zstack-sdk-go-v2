// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeFromVolumeBackupDetailParam CreateDataVolumeFromVolumeBackup detail param
type CreateDataVolumeFromVolumeBackupDetailParam struct {
	Name string `json:"name" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	BackupUuid string `json:"backupUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeBackupParam CreateDataVolumeFromVolumeBackup request param
type CreateDataVolumeFromVolumeBackupParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeBackupDetailParam `json:"params"`
}
