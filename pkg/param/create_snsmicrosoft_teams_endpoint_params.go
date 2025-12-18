// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSMicrosoftTeamsEndpointDetailParam CreateSNSMicrosoftTeamsEndpoint detail param
type CreateSNSMicrosoftTeamsEndpointDetailParam struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSMicrosoftTeamsEndpointParam CreateSNSMicrosoftTeamsEndpoint request param
type CreateSNSMicrosoftTeamsEndpointParam struct {
	BaseParam
	Params CreateSNSMicrosoftTeamsEndpointDetailParam `json:"params"`
}
