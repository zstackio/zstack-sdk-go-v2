// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSHttpEndpointDetailParam UpdateSNSHttpEndpoint detail param
type UpdateSNSHttpEndpointDetailParam struct {
	Url string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSHttpEndpointParam UpdateSNSHttpEndpoint request param
type UpdateSNSHttpEndpointParam struct {
	BaseParam
	Params UpdateSNSHttpEndpointDetailParam `json:"params"`
}
