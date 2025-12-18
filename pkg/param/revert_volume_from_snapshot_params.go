// Copyright (c) ZStack.io, Inc.

package param

// RevertVolumeFromSnapshotDetailParam RevertVolumeFromSnapshot detail param
type RevertVolumeFromSnapshotDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RevertVolumeFromSnapshotParam RevertVolumeFromSnapshot request param
type RevertVolumeFromSnapshotParam struct {
	BaseParam
	Params RevertVolumeFromSnapshotDetailParam `json:"params"`
}
