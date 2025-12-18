// Copyright (c) ZStack.io, Inc.

package param

// AddContainerManagementEndpointDetailParam AddContainerManagementEndpoint detail param
type AddContainerManagementEndpointDetailParam struct {
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
	Params AddContainerManagementEndpointDetailParam `json:"params"`
}
