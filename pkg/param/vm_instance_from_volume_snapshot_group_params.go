// Copyright (c) ZStack.io, Inc.

package param

// CreateVmInstanceFromVolumeSnapshotGroupDetailParam CreateVmInstanceFromVolumeSnapshotGroup详细参数
type CreateVmInstanceFromVolumeSnapshotGroupDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"instanceOfferingUuid,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest int64 `json:"reservedMemorySize,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"volumeSnapshotGroupUuid" validate:"required"` // 必填
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest map[string]interface{} `json:"dataVolumeSystemTags,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotGroupParam CreateVmInstanceFromVolumeSnapshotGroup请求参数
type CreateVmInstanceFromVolumeSnapshotGroupParam struct {
	BaseParam
	Params CreateVmInstanceFromVolumeSnapshotGroupDetailParam `json:"params"` // 详细参数
}

