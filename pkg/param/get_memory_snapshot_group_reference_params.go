// Copyright (c) ZStack.io, Inc.

package param

// GetMemorySnapshotGroupReferenceDetailParam GetMemorySnapshotGroupReference detail param
type GetMemorySnapshotGroupReferenceDetailParam struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	ResourceType string `json:"resourceType" validate:"required"`
}

// GetMemorySnapshotGroupReferenceParam GetMemorySnapshotGroupReference request param
type GetMemorySnapshotGroupReferenceParam struct {
	BaseParam
	Params GetMemorySnapshotGroupReferenceDetailParam `json:"params"`
}
