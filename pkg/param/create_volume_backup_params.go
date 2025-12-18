// Copyright (c) ZStack.io, Inc.

package param

// CreateVolumeBackupDetailParam CreateVolumeBackup detail param
type CreateVolumeBackupDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Mode string `json:"mode,omitempty"`
	VolumeReadBandwidth int64 `json:"volumeReadBandwidth,omitempty"`
	VolumeWriteBandwidth int64 `json:"volumeWriteBandwidth,omitempty"`
	NetworkReadBandwidth int64 `json:"networkReadBandwidth,omitempty"`
	NetworkWriteBandwidth int64 `json:"networkWriteBandwidth,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVolumeBackupParam CreateVolumeBackup request param
type CreateVolumeBackupParam struct {
	BaseParam
	Params CreateVolumeBackupDetailParam `json:"params"`
}
