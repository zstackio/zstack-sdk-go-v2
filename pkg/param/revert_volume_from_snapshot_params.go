// Copyright (c) ZStack.io, Inc.

package param

// RevertVolumeFromSnapshotDetailParam RevertVolumeFromSnapshot详细参数
type RevertVolumeFromSnapshotDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RevertVolumeFromSnapshotParam RevertVolumeFromSnapshot请求参数
type RevertVolumeFromSnapshotParam struct {
	BaseParam
	Params RevertVolumeFromSnapshotDetailParam `json:"params"` // 详细参数
}

