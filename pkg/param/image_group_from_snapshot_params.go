// Copyright (c) ZStack.io, Inc.

package param

// CreateImageGroupFromSnapshotDetailParam CreateImageGroupFromSnapshot详细参数
type CreateImageGroupFromSnapshotDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"rootVolumeSnapshotUuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest []string `json:"dataVolumeSnapshotUuids,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromSnapshotParam CreateImageGroupFromSnapshot请求参数
type CreateImageGroupFromSnapshotParam struct {
	BaseParam
	Params CreateImageGroupFromSnapshotDetailParam `json:"params"` // 详细参数
}

