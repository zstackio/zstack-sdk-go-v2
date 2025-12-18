// Copyright (c) ZStack.io, Inc.

package param

// UpdateContainerManagementEndpointDetailParam UpdateContainerManagementEndpoint detail param
type UpdateContainerManagementEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ManagementIp string `json:"managementIp,omitempty"`
	ManagementPort int `json:"managementPort,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

// UpdateContainerManagementEndpointParam UpdateContainerManagementEndpoint request param
type UpdateContainerManagementEndpointParam struct {
	BaseParam
	Params UpdateContainerManagementEndpointDetailParam `json:"params"`
}
