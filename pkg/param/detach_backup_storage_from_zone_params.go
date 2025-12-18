// Copyright (c) ZStack.io, Inc.

package param

// DetachBackupStorageFromZoneDetailParam DetachBackupStorageFromZone detail param
type DetachBackupStorageFromZoneDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
}

// DetachBackupStorageFromZoneParam DetachBackupStorageFromZone request param
type DetachBackupStorageFromZoneParam struct {
	BaseParam
	Params DetachBackupStorageFromZoneDetailParam `json:"params"`
}
