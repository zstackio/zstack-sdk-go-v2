// Copyright (c) ZStack.io, Inc.

package param

// DeleteContainerManagementEndpointDetailParam DeleteContainerManagementEndpoint detail param
type DeleteContainerManagementEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteContainerManagementEndpointParam DeleteContainerManagementEndpoint request param
type DeleteContainerManagementEndpointParam struct {
	BaseParam
	Params DeleteContainerManagementEndpointDetailParam `json:"params"`
}
