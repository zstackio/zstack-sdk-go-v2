// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateVolumeBackupParamDetail CreateVolumeBackup detail param
type CreateVolumeBackupParamDetail struct {
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
	Params CreateVolumeBackupParamDetail `json:"params"`
}
// SyncVolumeBackupParamDetail SyncVolumeBackup detail param
type SyncVolumeBackupParamDetail struct {
	ImageStoreUuid string `json:"imageStoreUuid" validate:"required"`
}

// SyncVolumeBackupParam SyncVolumeBackup request param
type SyncVolumeBackupParam struct {
	BaseParam
	Params SyncVolumeBackupParamDetail `json:"params"`
}
// DeleteVolumeBackupParamDetail DeleteVolumeBackup detail param
type DeleteVolumeBackupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	HandleDependency bool `json:"handleDependency,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVolumeBackupParam DeleteVolumeBackup request param
type DeleteVolumeBackupParam struct {
	BaseParam
	Params DeleteVolumeBackupParamDetail `json:"params"`
}
