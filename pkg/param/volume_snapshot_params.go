// Copyright (c) ZStack.io, Inc.

package param

// CreateVolumeSnapshotDetailParam CreateVolumeSnapshot详细参数
type CreateVolumeSnapshotDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVolumeSnapshotParam CreateVolumeSnapshot请求参数
type CreateVolumeSnapshotParam struct {
	BaseParam
	Params CreateVolumeSnapshotDetailParam `json:"params"` // 详细参数
}

