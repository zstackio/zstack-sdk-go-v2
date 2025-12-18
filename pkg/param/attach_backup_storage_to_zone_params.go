// Copyright (c) ZStack.io, Inc.

package param

// AttachBackupStorageToZoneDetailParam AttachBackupStorageToZone detail param
type AttachBackupStorageToZoneDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
}

// AttachBackupStorageToZoneParam AttachBackupStorageToZone request param
type AttachBackupStorageToZoneParam struct {
	BaseParam
	Params AttachBackupStorageToZoneDetailParam `json:"params"`
}
