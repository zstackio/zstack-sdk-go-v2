// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeFromVolumeSnapshotDetailParam CreateDataVolumeFromVolumeSnapshot detail param
type CreateDataVolumeFromVolumeSnapshotDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VolumeSnapshotUuid string `json:"volumeSnapshotUuid" validate:"required"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeSnapshotParam CreateDataVolumeFromVolumeSnapshot request param
type CreateDataVolumeFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeSnapshotDetailParam `json:"params"`
}
