// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteModelCenterParamDetail DeleteModelCenter detail param
type DeleteModelCenterParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteModelCenterParam DeleteModelCenter request param
type DeleteModelCenterParam struct {
	BaseParam
	DeleteModelCenter DeleteModelCenterParamDetail `json:"deleteModelCenter"`
}
// AddModelCenterParamDetail AddModelCenter detail param
type AddModelCenterParamDetail struct {
	Name string `json:"name" validate:"required"`
	Url string `json:"url" validate:"required"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	ManagementPort int `json:"managementPort" validate:"required"`
	Parameters string `json:"parameters,omitempty"`
	StorageNetworkUuid string `json:"storageNetworkUuid,omitempty"`
	ServiceNetworkUuid string `json:"serviceNetworkUuid,omitempty"`
	ContainerRegistry string `json:"containerRegistry,omitempty"`
	ContainerNetwork string `json:"containerNetwork,omitempty"`
	ContainerStorageNetwork string `json:"containerStorageNetwork,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddModelCenterParam AddModelCenter request param
type AddModelCenterParam struct {
	BaseParam
	AddModelCenter AddModelCenterParamDetail `json:"addModelCenter"`
}
// UpdateModelCenterParamDetail UpdateModelCenter detail param
type UpdateModelCenterParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	ManagementPort int `json:"managementPort,omitempty"`
	StorageNetworkUuid string `json:"storageNetworkUuid,omitempty"`
	ServiceNetworkUuid string `json:"serviceNetworkUuid,omitempty"`
	ContainerRegistry string `json:"containerRegistry,omitempty"`
	ContainerNetwork string `json:"containerNetwork,omitempty"`
	ContainerStorageNetwork string `json:"containerStorageNetwork,omitempty"`
}

// UpdateModelCenterParam UpdateModelCenter request param
type UpdateModelCenterParam struct {
	BaseParam
	UpdateModelCenter UpdateModelCenterParamDetail `json:"updateModelCenter"`
}
