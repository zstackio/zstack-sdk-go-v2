// Copyright (c) ZStack.io, Inc.

package param

// CreateVmInstanceFromVolumeSnapshotGroupDetailParam CreateVmInstanceFromVolumeSnapshotGroup detail param
type CreateVmInstanceFromVolumeSnapshotGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	Type string `json:"type,omitempty"`
	VolumeSnapshotGroupUuid string `json:"volumeSnapshotGroupUuid" validate:"required"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags map[string]interface{} `json:"dataVolumeSystemTags,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotGroupParam CreateVmInstanceFromVolumeSnapshotGroup request param
type CreateVmInstanceFromVolumeSnapshotGroupParam struct {
	BaseParam
	Params CreateVmInstanceFromVolumeSnapshotGroupDetailParam `json:"params"`
}
