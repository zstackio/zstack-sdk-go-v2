// Copyright (c) ZStack.io, Inc.

package param

// CreateVmBackupDetailParam CreateVmBackup detail param
type CreateVmBackupDetailParam struct {
	RootVolumeUuid string `json:"rootVolumeUuid" validate:"required"`
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

// CreateVmBackupParam CreateVmBackup request param
type CreateVmBackupParam struct {
	BaseParam
	Params CreateVmBackupDetailParam `json:"params"`
}
