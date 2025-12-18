// Copyright (c) ZStack.io, Inc.

package param

// CreateVolumesSnapshotDetailParam CreateVolumesSnapshot detail param
type CreateVolumesSnapshotDetailParam struct {
	VolumeUuids []string `json:"volumeUuids" validate:"required"`
}

// CreateVolumesSnapshotParam CreateVolumesSnapshot request param
type CreateVolumesSnapshotParam struct {
	BaseParam
	Params CreateVolumesSnapshotDetailParam `json:"params"`
}
