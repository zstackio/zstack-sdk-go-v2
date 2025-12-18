// Copyright (c) ZStack.io, Inc.

package param

// DetachBackupStorageFromZoneDetailParam DetachBackupStorageFromZone详细参数
type DetachBackupStorageFromZoneDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"zoneUuid" validate:"required"` // 必填
}

// DetachBackupStorageFromZoneParam DetachBackupStorageFromZone请求参数
type DetachBackupStorageFromZoneParam struct {
	BaseParam
	Params DetachBackupStorageFromZoneDetailParam `json:"params"` // 详细参数
}

