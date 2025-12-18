// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeTemplateFromVolumeSnapshotDetailParam CreateDataVolumeTemplateFromVolumeSnapshot详细参数
type CreateDataVolumeTemplateFromVolumeSnapshotDetailParam struct {
	rest string `json:"snapshotUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest []string `json:"backupStorageUuids" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeSnapshotParam CreateDataVolumeTemplateFromVolumeSnapshot请求参数
type CreateDataVolumeTemplateFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateDataVolumeTemplateFromVolumeSnapshotDetailParam `json:"params"` // 详细参数
}

