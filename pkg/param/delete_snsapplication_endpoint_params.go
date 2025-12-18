// Copyright (c) ZStack.io, Inc.

package param

// DeleteSNSApplicationEndpointDetailParam DeleteSNSApplicationEndpoint detail param
type DeleteSNSApplicationEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSNSApplicationEndpointParam DeleteSNSApplicationEndpoint request param
type DeleteSNSApplicationEndpointParam struct {
	BaseParam
	Params DeleteSNSApplicationEndpointDetailParam `json:"params"`
}
