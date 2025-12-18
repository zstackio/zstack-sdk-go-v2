// Copyright (c) ZStack.io, Inc.

package param

// CreateBareMetal2InstanceDetailParam CreateBareMetal2Instance detail param
type CreateBareMetal2InstanceDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	GatewayUuid string `json:"gatewayUuid,omitempty"`
	ChassisUuid string `json:"chassisUuid,omitempty"`
	ChassisOfferingUuid string `json:"chassisOfferingUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	ChassisDiskUuid string `json:"chassisDiskUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume string `json:"primaryStorageUuidForDataVolume,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	GatewayAllocatorStrategy string `json:"gatewayAllocatorStrategy,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBareMetal2InstanceParam CreateBareMetal2Instance request param
type CreateBareMetal2InstanceParam struct {
	BaseParam
	Params CreateBareMetal2InstanceDetailParam `json:"params"`
}
