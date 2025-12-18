// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsImageFromEcsSnapshotDetailParam CreateEcsImageFromEcsSnapshot detail param
type CreateEcsImageFromEcsSnapshotDetailParam struct {
	SnapshotUuid string `json:"snapshotUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateEcsImageFromEcsSnapshotParam CreateEcsImageFromEcsSnapshot request param
type CreateEcsImageFromEcsSnapshotParam struct {
	BaseParam
	Params CreateEcsImageFromEcsSnapshotDetailParam `json:"params"`
}
