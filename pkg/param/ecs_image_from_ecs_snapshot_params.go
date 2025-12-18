// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsImageFromEcsSnapshotDetailParam CreateEcsImageFromEcsSnapshot详细参数
type CreateEcsImageFromEcsSnapshotDetailParam struct {
	rest string `json:"snapshotUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateEcsImageFromEcsSnapshotParam CreateEcsImageFromEcsSnapshot请求参数
type CreateEcsImageFromEcsSnapshotParam struct {
	BaseParam
	Params CreateEcsImageFromEcsSnapshotDetailParam `json:"params"` // 详细参数
}

