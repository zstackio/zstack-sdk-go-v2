// Copyright (c) ZStack.io, Inc.

package param

// SNSEmailTestConnectionDetailParam SNSEmailTestConnection detail param
type SNSEmailTestConnectionDetailParam struct {
	Emails []string `json:"emails,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	Subject string `json:"subject,omitempty"`
	Text string `json:"text,omitempty"`
}

// SNSEmailTestConnectionParam SNSEmailTestConnection request param
type SNSEmailTestConnectionParam struct {
	BaseParam
	Params SNSEmailTestConnectionDetailParam `json:"params"`
}
