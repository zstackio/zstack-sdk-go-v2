// Copyright (c) ZStack.io, Inc.

package param

// DeleteContainerResourceFromEndpointDetailParam DeleteContainerResourceFromEndpoint detail param
type DeleteContainerResourceFromEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteContainerResourceFromEndpointParam DeleteContainerResourceFromEndpoint request param
type DeleteContainerResourceFromEndpointParam struct {
	BaseParam
	Params DeleteContainerResourceFromEndpointDetailParam `json:"params"`
}
