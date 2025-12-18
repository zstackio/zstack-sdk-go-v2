// Copyright (c) ZStack.io, Inc.

package param

// AddModelCenterDetailParam AddModelCenter detail param
type AddModelCenterDetailParam struct {
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
	Params AddModelCenterDetailParam `json:"params"`
}
