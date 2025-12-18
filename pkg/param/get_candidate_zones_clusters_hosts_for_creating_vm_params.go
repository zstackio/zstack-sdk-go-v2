// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateZonesClustersHostsForCreatingVmDetailParam GetCandidateZonesClustersHostsForCreatingVm detail param
type GetCandidateZonesClustersHostsForCreatingVmDetailParam struct {
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	RootDiskOfferingUuid string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize int64 `json:"rootDiskSize,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
}

// GetCandidateZonesClustersHostsForCreatingVmParam GetCandidateZonesClustersHostsForCreatingVm request param
type GetCandidateZonesClustersHostsForCreatingVmParam struct {
	BaseParam
	Params GetCandidateZonesClustersHostsForCreatingVmDetailParam `json:"params"`
}
