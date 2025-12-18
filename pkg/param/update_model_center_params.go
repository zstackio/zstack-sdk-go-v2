// Copyright (c) ZStack.io, Inc.

package param

// UpdateModelCenterDetailParam UpdateModelCenter detail param
type UpdateModelCenterDetailParam struct {
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
	Params UpdateModelCenterDetailParam `json:"params"`
}
