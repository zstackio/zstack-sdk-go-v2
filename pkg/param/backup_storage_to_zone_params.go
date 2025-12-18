// Copyright (c) ZStack.io, Inc.

package param

// AttachBackupStorageToZoneDetailParam AttachBackupStorageToZone详细参数
type AttachBackupStorageToZoneDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
}

// AttachBackupStorageToZoneParam AttachBackupStorageToZone请求参数
type AttachBackupStorageToZoneParam struct {
	BaseParam
	Params AttachBackupStorageToZoneDetailParam `json:"params"` // 详细参数
}

