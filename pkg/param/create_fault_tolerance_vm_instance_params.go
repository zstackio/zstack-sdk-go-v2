// Copyright (c) ZStack.io, Inc.

package param

// CreateFaultToleranceVmInstanceDetailParam CreateFaultToleranceVmInstance detail param
type CreateFaultToleranceVmInstanceDetailParam struct {
	Name string `json:"name" validate:"required"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	Description string `json:"description,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFaultToleranceVmInstanceParam CreateFaultToleranceVmInstance request param
type CreateFaultToleranceVmInstanceParam struct {
	BaseParam
	Params CreateFaultToleranceVmInstanceDetailParam `json:"params"`
}
