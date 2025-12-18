// Copyright (c) ZStack.io, Inc.

package param

// ShrinkVolumeSnapshotDetailParam ShrinkVolumeSnapshot详细参数
type ShrinkVolumeSnapshotDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ShrinkVolumeSnapshotParam ShrinkVolumeSnapshot请求参数
type ShrinkVolumeSnapshotParam struct {
	BaseParam
	Params ShrinkVolumeSnapshotDetailParam `json:"params"` // 详细参数
}

