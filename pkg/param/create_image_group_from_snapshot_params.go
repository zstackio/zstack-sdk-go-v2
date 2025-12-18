// Copyright (c) ZStack.io, Inc.

package param

// CreateImageGroupFromSnapshotDetailParam CreateImageGroupFromSnapshot detail param
type CreateImageGroupFromSnapshotDetailParam struct {
	Name string `json:"name" validate:"required"`
	RootVolumeSnapshotUuid string `json:"rootVolumeSnapshotUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	DataVolumeSnapshotUuids []string `json:"dataVolumeSnapshotUuids,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromSnapshotParam CreateImageGroupFromSnapshot request param
type CreateImageGroupFromSnapshotParam struct {
	BaseParam
	Params CreateImageGroupFromSnapshotDetailParam `json:"params"`
}
