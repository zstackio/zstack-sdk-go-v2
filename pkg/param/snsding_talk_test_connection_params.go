// Copyright (c) ZStack.io, Inc.

package param

// SNSDingTalkTestConnectionDetailParam SNSDingTalkTestConnection detail param
type SNSDingTalkTestConnectionDetailParam struct {
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonPhoneNumbers []string `json:"atPersonPhoneNumbers,omitempty"`
	Secret string `json:"secret,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// SNSDingTalkTestConnectionParam SNSDingTalkTestConnection request param
type SNSDingTalkTestConnectionParam struct {
	BaseParam
	Params SNSDingTalkTestConnectionDetailParam `json:"params"`
}
