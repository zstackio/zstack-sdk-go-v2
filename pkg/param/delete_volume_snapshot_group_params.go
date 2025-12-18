// Copyright (c) ZStack.io, Inc.

package param

// DeleteVolumeSnapshotGroupDetailParam DeleteVolumeSnapshotGroup detail param
type DeleteVolumeSnapshotGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Direction string `json:"direction,omitempty"`
	Scope string `json:"scope,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVolumeSnapshotGroupParam DeleteVolumeSnapshotGroup request param
type DeleteVolumeSnapshotGroupParam struct {
	BaseParam
	Params DeleteVolumeSnapshotGroupDetailParam `json:"params"`
}
