// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSWeComEndpointDetailParam UpdateSNSWeComEndpoint detail param
type UpdateSNSWeComEndpointDetailParam struct {
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSWeComEndpointParam UpdateSNSWeComEndpoint request param
type UpdateSNSWeComEndpointParam struct {
	BaseParam
	Params UpdateSNSWeComEndpointDetailParam `json:"params"`
}
