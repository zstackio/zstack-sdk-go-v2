// Copyright (c) ZStack.io, Inc.

package param

// CheckVolumeSnapshotGroupAvailabilityDetailParam CheckVolumeSnapshotGroupAvailability详细参数
type CheckVolumeSnapshotGroupAvailabilityDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
}

// CheckVolumeSnapshotGroupAvailabilityParam CheckVolumeSnapshotGroupAvailability请求参数
type CheckVolumeSnapshotGroupAvailabilityParam struct {
	BaseParam
	Params CheckVolumeSnapshotGroupAvailabilityDetailParam `json:"params"` // 详细参数
}

