// Copyright (c) ZStack.io, Inc.

package param

// UngroupVolumeSnapshotGroupDetailParam UngroupVolumeSnapshotGroup detail param
type UngroupVolumeSnapshotGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// UngroupVolumeSnapshotGroupParam UngroupVolumeSnapshotGroup request param
type UngroupVolumeSnapshotGroupParam struct {
	BaseParam
	Params UngroupVolumeSnapshotGroupDetailParam `json:"params"`
}
