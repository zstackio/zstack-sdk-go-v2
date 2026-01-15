// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateContainerManagementEndpointParamDetail UpdateContainerManagementEndpoint detail param
type UpdateContainerManagementEndpointParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	ManagementIp string `json:"managementIp,omitempty"`
	ManagementPort int `json:"managementPort,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

// UpdateContainerManagementEndpointParam UpdateContainerManagementEndpoint request param
type UpdateContainerManagementEndpointParam struct {
	BaseParam
	UpdateContainerManagementEndpoint UpdateContainerManagementEndpointParamDetail `json:"updateContainerManagementEndpoint"`
}
// AddContainerManagementEndpointParamDetail AddContainerManagementEndpoint detail param
type AddContainerManagementEndpointParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	Vendor string `json:"vendor" validate:"required"`
	ManagementPort int `json:"managementPort" validate:"required"`
	ContainerAccessKeyId string `json:"containerAccessKeyId" validate:"required"`
	ContainerAccessKeySecret string `json:"containerAccessKeySecret" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddContainerManagementEndpointParam AddContainerManagementEndpoint request param
type AddContainerManagementEndpointParam struct {
	BaseParam
	AddContainerManagementEndpoint AddContainerManagementEndpointParamDetail `json:"addContainerManagementEndpoint"`
}
// SyncContainerManagementEndpointParamDetail SyncContainerManagementEndpoint detail param
type SyncContainerManagementEndpointParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
}

// SyncContainerManagementEndpointParam SyncContainerManagementEndpoint request param
type SyncContainerManagementEndpointParam struct {
	BaseParam
	SyncContainerManagementEndpoint SyncContainerManagementEndpointParamDetail `json:"syncContainerManagementEndpoint"`
}
// DeleteContainerManagementEndpointParamDetail DeleteContainerManagementEndpoint detail param
type DeleteContainerManagementEndpointParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteContainerManagementEndpointParam DeleteContainerManagementEndpoint request param
type DeleteContainerManagementEndpointParam struct {
	BaseParam
	DeleteContainerManagementEndpoint DeleteContainerManagementEndpointParamDetail `json:"deleteContainerManagementEndpoint"`
}
