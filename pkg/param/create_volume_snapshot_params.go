// Copyright (c) ZStack.io, Inc.

package param

// CreateVolumeSnapshotDetailParam CreateVolumeSnapshot detail param
type CreateVolumeSnapshotDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVolumeSnapshotParam CreateVolumeSnapshot request param
type CreateVolumeSnapshotParam struct {
	BaseParam
	Params CreateVolumeSnapshotDetailParam `json:"params"`
}
