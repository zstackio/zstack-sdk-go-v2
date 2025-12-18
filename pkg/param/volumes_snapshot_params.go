// Copyright (c) ZStack.io, Inc.

package param

// CreateVolumesSnapshotDetailParam CreateVolumesSnapshot详细参数
type CreateVolumesSnapshotDetailParam struct {
	rest []string `json:"volumeUuids" validate:"required"` // 必填
}

// CreateVolumesSnapshotParam CreateVolumesSnapshot请求参数
type CreateVolumesSnapshotParam struct {
	BaseParam
	Params CreateVolumesSnapshotDetailParam `json:"params"` // 详细参数
}

