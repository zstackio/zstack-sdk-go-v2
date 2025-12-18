// Copyright (c) ZStack.io, Inc.

package param

// BatchDeleteVolumeSnapshotDetailParam BatchDeleteVolumeSnapshot详细参数
type BatchDeleteVolumeSnapshotDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// BatchDeleteVolumeSnapshotParam BatchDeleteVolumeSnapshot请求参数
type BatchDeleteVolumeSnapshotParam struct {
	BaseParam
	Params BatchDeleteVolumeSnapshotDetailParam `json:"params"` // 详细参数
}

