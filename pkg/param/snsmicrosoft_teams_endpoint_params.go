// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSMicrosoftTeamsEndpointDetailParam UpdateSNSMicrosoftTeamsEndpoint详细参数
type UpdateSNSMicrosoftTeamsEndpointDetailParam struct {
	rest string `json:"url,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"platformUuid,omitempty"`
}

// UpdateSNSMicrosoftTeamsEndpointParam UpdateSNSMicrosoftTeamsEndpoint请求参数
type UpdateSNSMicrosoftTeamsEndpointParam struct {
	BaseParam
	Params UpdateSNSMicrosoftTeamsEndpointDetailParam `json:"params"` // 详细参数
}

