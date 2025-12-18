// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSMicrosoftTeamsEndpointDetailParam UpdateSNSMicrosoftTeamsEndpoint detail param
type UpdateSNSMicrosoftTeamsEndpointDetailParam struct {
	Url string `json:"url,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSMicrosoftTeamsEndpointParam UpdateSNSMicrosoftTeamsEndpoint request param
type UpdateSNSMicrosoftTeamsEndpointParam struct {
	BaseParam
	Params UpdateSNSMicrosoftTeamsEndpointDetailParam `json:"params"`
}
