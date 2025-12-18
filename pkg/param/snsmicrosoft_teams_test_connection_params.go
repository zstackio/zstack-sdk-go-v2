// Copyright (c) ZStack.io, Inc.

package param

// SNSMicrosoftTeamsTestConnectionDetailParam SNSMicrosoftTeamsTestConnection detail param
type SNSMicrosoftTeamsTestConnectionDetailParam struct {
	Url string `json:"url,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// SNSMicrosoftTeamsTestConnectionParam SNSMicrosoftTeamsTestConnection request param
type SNSMicrosoftTeamsTestConnectionParam struct {
	BaseParam
	Params SNSMicrosoftTeamsTestConnectionDetailParam `json:"params"`
}
