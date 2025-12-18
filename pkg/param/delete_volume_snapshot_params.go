// Copyright (c) ZStack.io, Inc.

package param

// DeleteVolumeSnapshotDetailParam DeleteVolumeSnapshot detail param
type DeleteVolumeSnapshotDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Direction string `json:"direction,omitempty"`
	Scope string `json:"scope,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVolumeSnapshotParam DeleteVolumeSnapshot request param
type DeleteVolumeSnapshotParam struct {
	BaseParam
	Params DeleteVolumeSnapshotDetailParam `json:"params"`
}
