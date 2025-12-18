// Copyright (c) ZStack.io, Inc.

package param

// CheckVolumeSnapshotGroupAvailabilityDetailParam CheckVolumeSnapshotGroupAvailability detail param
type CheckVolumeSnapshotGroupAvailabilityDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
}

// CheckVolumeSnapshotGroupAvailabilityParam CheckVolumeSnapshotGroupAvailability request param
type CheckVolumeSnapshotGroupAvailabilityParam struct {
	BaseParam
	Params CheckVolumeSnapshotGroupAvailabilityDetailParam `json:"params"`
}
