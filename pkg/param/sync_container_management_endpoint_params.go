// Copyright (c) ZStack.io, Inc.

package param

// SyncContainerManagementEndpointDetailParam SyncContainerManagementEndpoint detail param
type SyncContainerManagementEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
}

// SyncContainerManagementEndpointParam SyncContainerManagementEndpoint request param
type SyncContainerManagementEndpointParam struct {
	BaseParam
	Params SyncContainerManagementEndpointDetailParam `json:"params"`
}
