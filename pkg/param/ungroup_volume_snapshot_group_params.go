// Copyright (c) ZStack.io, Inc.

package param

// UngroupVolumeSnapshotGroupDetailParam UngroupVolumeSnapshotGroup详细参数
type UngroupVolumeSnapshotGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// UngroupVolumeSnapshotGroupParam UngroupVolumeSnapshotGroup请求参数
type UngroupVolumeSnapshotGroupParam struct {
	BaseParam
	Params UngroupVolumeSnapshotGroupDetailParam `json:"params"` // 详细参数
}

