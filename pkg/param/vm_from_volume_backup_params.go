// Copyright (c) ZStack.io, Inc.

package param

// CreateVmFromVolumeBackupDetailParam CreateVmFromVolumeBackup详细参数
type CreateVmFromVolumeBackupDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"backupUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"instanceOfferingUuid" validate:"required"` // 必填
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest []string `json:"l3NetworkUuids" validate:"required"` // 必填
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"description,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVmFromVolumeBackupParam CreateVmFromVolumeBackup请求参数
type CreateVmFromVolumeBackupParam struct {
	BaseParam
	Params CreateVmFromVolumeBackupDetailParam `json:"params"` // 详细参数
}

