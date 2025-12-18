// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeFromVolumeSnapshotDetailParam CreateDataVolumeFromVolumeSnapshot详细参数
type CreateDataVolumeFromVolumeSnapshotDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"volumeSnapshotUuid" validate:"required"` // 必填
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeSnapshotParam CreateDataVolumeFromVolumeSnapshot请求参数
type CreateDataVolumeFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeSnapshotDetailParam `json:"params"` // 详细参数
}

