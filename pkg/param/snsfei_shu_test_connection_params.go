// Copyright (c) ZStack.io, Inc.

package param

// SNSFeiShuTestConnectionDetailParam SNSFeiShuTestConnection detail param
type SNSFeiShuTestConnectionDetailParam struct {
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	Secret string `json:"secret,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// SNSFeiShuTestConnectionParam SNSFeiShuTestConnection request param
type SNSFeiShuTestConnectionParam struct {
	BaseParam
	Params SNSFeiShuTestConnectionDetailParam `json:"params"`
}
