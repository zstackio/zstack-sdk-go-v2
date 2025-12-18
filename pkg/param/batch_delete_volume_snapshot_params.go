// Copyright (c) ZStack.io, Inc.

package param

// BatchDeleteVolumeSnapshotDetailParam BatchDeleteVolumeSnapshot detail param
type BatchDeleteVolumeSnapshotDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// BatchDeleteVolumeSnapshotParam BatchDeleteVolumeSnapshot request param
type BatchDeleteVolumeSnapshotParam struct {
	BaseParam
	Params BatchDeleteVolumeSnapshotDetailParam `json:"params"`
}
