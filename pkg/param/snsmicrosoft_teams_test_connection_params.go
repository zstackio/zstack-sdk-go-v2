// Copyright (c) ZStack.io, Inc.

package param

// SNSMicrosoftTeamsTestConnectionDetailParam SNSMicrosoftTeamsTestConnection详细参数
type SNSMicrosoftTeamsTestConnectionDetailParam struct {
	rest string `json:"url,omitempty"`
	rest string `json:"testMsg" validate:"required"` // 必填
	rest string `json:"endpointUuid,omitempty"`
}

// SNSMicrosoftTeamsTestConnectionParam SNSMicrosoftTeamsTestConnection请求参数
type SNSMicrosoftTeamsTestConnectionParam struct {
	BaseParam
	Params SNSMicrosoftTeamsTestConnectionDetailParam `json:"params"` // 详细参数
}

