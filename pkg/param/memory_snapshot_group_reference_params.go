// Copyright (c) ZStack.io, Inc.

package param

// GetMemorySnapshotGroupReferenceDetailParam GetMemorySnapshotGroupReference详细参数
type GetMemorySnapshotGroupReferenceDetailParam struct {
	rest string `json:"resourceUuid" validate:"required"` // 必填
	rest string `json:"resourceType" validate:"required"` // 必填
}

// GetMemorySnapshotGroupReferenceParam GetMemorySnapshotGroupReference请求参数
type GetMemorySnapshotGroupReferenceParam struct {
	BaseParam
	Params GetMemorySnapshotGroupReferenceDetailParam `json:"params"` // 详细参数
}

