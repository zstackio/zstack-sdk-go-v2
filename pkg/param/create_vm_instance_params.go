// Copyright (c) ZStack.io, Inc.

package param

// CreateVmInstanceDetailParam CreateVmInstance detail param
type CreateVmInstanceDetailParam struct {
	Name string `json:"name" validate:"required"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams string `json:"vmNicParams,omitempty"`
	Type string `json:"type,omitempty"`
	RootDiskOfferingUuid string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize int64 `json:"rootDiskSize,omitempty"`
	DataDiskSizes []int64 `json:"dataDiskSizes,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	Description string `json:"description,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	DataVolumeSystemTagsOnIndex map[string]interface{} `json:"dataVolumeSystemTagsOnIndex,omitempty"`
	SshKeyPairUuids []string `json:"sshKeyPairUuids,omitempty"`
	Platform string `json:"platform,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	DiskAOs []DiskAOParam `json:"diskAOs,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceParam CreateVmInstance request param
type CreateVmInstanceParam struct {
	BaseParam
	Params CreateVmInstanceDetailParam `json:"params"`
}
