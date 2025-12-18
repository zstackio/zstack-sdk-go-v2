// Copyright (c) ZStack.io, Inc.

package param

// CreateVmFromCdpBackupDetailParam CreateVmFromCdpBackup详细参数
type CreateVmFromCdpBackupDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest int64 `json:"groupId" validate:"required"` // 必填
	rest string `json:"cdpTaskUuid" validate:"required"` // 必填
	rest string `json:"instanceOfferingUuid" validate:"required"` // 必填
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"primaryStorageUuidForDataVolume,omitempty"`
	rest int64 `json:"recoverBandwidth,omitempty"`
	rest string `json:"description,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest []string `json:"dataVolumeSystemTags,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVmFromCdpBackupParam CreateVmFromCdpBackup请求参数
type CreateVmFromCdpBackupParam struct {
	BaseParam
	Params CreateVmFromCdpBackupDetailParam `json:"params"` // 详细参数
}

