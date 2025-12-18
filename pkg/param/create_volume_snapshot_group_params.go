// Copyright (c) ZStack.io, Inc.

package param

// CreateVolumeSnapshotGroupDetailParam CreateVolumeSnapshotGroup detail param
type CreateVolumeSnapshotGroupDetailParam struct {
	RootVolumeUuid string `json:"rootVolumeUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	WithMemory bool `json:"withMemory,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVolumeSnapshotGroupParam CreateVolumeSnapshotGroup request param
type CreateVolumeSnapshotGroupParam struct {
	BaseParam
	Params CreateVolumeSnapshotGroupDetailParam `json:"params"`
}
