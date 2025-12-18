// Copyright (c) ZStack.io, Inc.

package param

// CreateVmInstanceFromVolumeSnapshotDetailParam CreateVmInstanceFromVolumeSnapshot详细参数
type CreateVmInstanceFromVolumeSnapshotDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest int64 `json:"reservedMemorySize,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest string `json:"vmNicParams,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"volumeSnapshotUuid" validate:"required"` // 必填
	rest string `json:"platform,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotParam CreateVmInstanceFromVolumeSnapshot请求参数
type CreateVmInstanceFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateVmInstanceFromVolumeSnapshotDetailParam `json:"params"` // 详细参数
}

