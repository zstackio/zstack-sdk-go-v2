// Copyright (c) ZStack.io, Inc.

package param

// GetCandidatePrimaryStoragesForCreatingVmDetailParam GetCandidatePrimaryStoragesForCreatingVm detail param
type GetCandidatePrimaryStoragesForCreatingVmDetailParam struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	RootDiskOfferingUuid string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize int64 `json:"rootDiskSize,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	DataDiskSizes []int64 `json:"dataDiskSizes,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
}

// GetCandidatePrimaryStoragesForCreatingVmParam GetCandidatePrimaryStoragesForCreatingVm request param
type GetCandidatePrimaryStoragesForCreatingVmParam struct {
	BaseParam
	Params GetCandidatePrimaryStoragesForCreatingVmDetailParam `json:"params"`
}
