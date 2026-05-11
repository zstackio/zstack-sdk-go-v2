// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateBareMetal2InstanceParamDetail CreateBareMetal2Instance detail param
type CreateBareMetal2InstanceParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	GatewayUuid *string `json:"gatewayUuid,omitempty"`
	ChassisUuid *string `json:"chassisUuid,omitempty"`
	ChassisOfferingUuid *string `json:"chassisOfferingUuid,omitempty"`
	ImageUuid *string `json:"imageUuid,omitempty"`
	ChassisDiskUuid *string `json:"chassisDiskUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume *string `json:"primaryStorageUuidForDataVolume,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	GatewayAllocatorStrategy *string `json:"gatewayAllocatorStrategy,omitempty"`
	ChassisType *string `json:"chassisType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBareMetal2InstanceParam CreateBareMetal2Instance request param
type CreateBareMetal2InstanceParam struct {
	BaseParam
	Params CreateBareMetal2InstanceParamDetail `json:"params"`
}
// ReconnectBareMetal2InstanceParamDetail ReconnectBareMetal2Instance detail param
type ReconnectBareMetal2InstanceParamDetail struct {
}

// ReconnectBareMetal2InstanceParam ReconnectBareMetal2Instance request param
type ReconnectBareMetal2InstanceParam struct {
	BaseParam
	Params ReconnectBareMetal2InstanceParamDetail `json:"reconnectBareMetal2Instance"`
}
// StartBareMetal2InstanceParamDetail StartBareMetal2Instance detail param
type StartBareMetal2InstanceParamDetail struct {
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	GatewayUuid *string `json:"gatewayUuid,omitempty"`
	ChassisUuid *string `json:"chassisUuid,omitempty"`
	ChassisOfferingUuid *string `json:"chassisOfferingUuid,omitempty"`
	ChassisType *string `json:"chassisType,omitempty"`
}

// StartBareMetal2InstanceParam StartBareMetal2Instance request param
type StartBareMetal2InstanceParam struct {
	BaseParam
	Params StartBareMetal2InstanceParamDetail `json:"startBareMetal2Instance"`
}
// UpdateBareMetal2InstanceParamDetail UpdateBareMetal2Instance detail param
type UpdateBareMetal2InstanceParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	ChassisOfferingUuid *string `json:"chassisOfferingUuid,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	AutoReleaseChassisEvent *string `json:"autoReleaseChassisEvent,omitempty"`
}

// UpdateBareMetal2InstanceParam UpdateBareMetal2Instance request param
type UpdateBareMetal2InstanceParam struct {
	BaseParam
	Params UpdateBareMetal2InstanceParamDetail `json:"updateBareMetal2Instance"`
}
