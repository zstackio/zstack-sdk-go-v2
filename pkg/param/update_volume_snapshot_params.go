// Copyright (c) ZStack.io, Inc.

package param

// UpdateVolumeSnapshotDetailParam UpdateVolumeSnapshot detail param
type UpdateVolumeSnapshotDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVolumeSnapshotParam UpdateVolumeSnapshot request param
type UpdateVolumeSnapshotParam struct {
	BaseParam
	Params UpdateVolumeSnapshotDetailParam `json:"params"`
}
