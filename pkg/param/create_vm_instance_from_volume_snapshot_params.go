// Copyright (c) ZStack.io, Inc.

package param

// CreateVmInstanceFromVolumeSnapshotDetailParam CreateVmInstanceFromVolumeSnapshot detail param
type CreateVmInstanceFromVolumeSnapshotDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams string `json:"vmNicParams,omitempty"`
	Type string `json:"type,omitempty"`
	VolumeSnapshotUuid string `json:"volumeSnapshotUuid" validate:"required"`
	Platform string `json:"platform,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotParam CreateVmInstanceFromVolumeSnapshot request param
type CreateVmInstanceFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateVmInstanceFromVolumeSnapshotDetailParam `json:"params"`
}
