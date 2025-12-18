// Copyright (c) ZStack.io, Inc.

package param

// UpdateVolumeSnapshotGroupDetailParam UpdateVolumeSnapshotGroup detail param
type UpdateVolumeSnapshotGroupDetailParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
}

// UpdateVolumeSnapshotGroupParam UpdateVolumeSnapshotGroup request param
type UpdateVolumeSnapshotGroupParam struct {
	BaseParam
	Params UpdateVolumeSnapshotGroupDetailParam `json:"params"`
}
