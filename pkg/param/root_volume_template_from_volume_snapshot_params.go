// Copyright (c) ZStack.io, Inc.

package param

// CreateRootVolumeTemplateFromVolumeSnapshotDetailParam CreateRootVolumeTemplateFromVolumeSnapshot详细参数
type CreateRootVolumeTemplateFromVolumeSnapshotDetailParam struct {
	rest string `json:"snapshotUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"guestOsType,omitempty"`
	rest []string `json:"backupStorageUuids" validate:"required"` // 必填
	rest string `json:"platform,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest bool `json:"system,omitempty"`
	rest bool `json:"virtio,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeSnapshotParam CreateRootVolumeTemplateFromVolumeSnapshot请求参数
type CreateRootVolumeTemplateFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromVolumeSnapshotDetailParam `json:"params"` // 详细参数
}

