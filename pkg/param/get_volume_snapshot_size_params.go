// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeSnapshotSizeDetailParam GetVolumeSnapshotSize detail param
type GetVolumeSnapshotSizeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVolumeSnapshotSizeParam GetVolumeSnapshotSize request param
type GetVolumeSnapshotSizeParam struct {
	BaseParam
	Params GetVolumeSnapshotSizeDetailParam `json:"params"`
}
