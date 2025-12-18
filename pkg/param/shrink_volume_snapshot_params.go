// Copyright (c) ZStack.io, Inc.

package param

// ShrinkVolumeSnapshotDetailParam ShrinkVolumeSnapshot detail param
type ShrinkVolumeSnapshotDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ShrinkVolumeSnapshotParam ShrinkVolumeSnapshot request param
type ShrinkVolumeSnapshotParam struct {
	BaseParam
	Params ShrinkVolumeSnapshotDetailParam `json:"params"`
}
