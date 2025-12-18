// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeSnapshotSizeDetailParam GetVolumeSnapshotSize详细参数
type GetVolumeSnapshotSizeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVolumeSnapshotSizeParam GetVolumeSnapshotSize请求参数
type GetVolumeSnapshotSizeParam struct {
	BaseParam
	Params GetVolumeSnapshotSizeDetailParam `json:"params"` // 详细参数
}

