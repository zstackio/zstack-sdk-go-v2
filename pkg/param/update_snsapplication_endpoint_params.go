// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSApplicationEndpointDetailParam UpdateSNSApplicationEndpoint detail param
type UpdateSNSApplicationEndpointDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSApplicationEndpointParam UpdateSNSApplicationEndpoint request param
type UpdateSNSApplicationEndpointParam struct {
	BaseParam
	Params UpdateSNSApplicationEndpointDetailParam `json:"params"`
}
