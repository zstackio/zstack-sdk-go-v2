// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSFeiShuEndpointDetailParam UpdateSNSFeiShuEndpoint detail param
type UpdateSNSFeiShuEndpointDetailParam struct {
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	Secret string `json:"secret,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSFeiShuEndpointParam UpdateSNSFeiShuEndpoint request param
type UpdateSNSFeiShuEndpointParam struct {
	BaseParam
	Params UpdateSNSFeiShuEndpointDetailParam `json:"params"`
}
