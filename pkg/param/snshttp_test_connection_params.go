// Copyright (c) ZStack.io, Inc.

package param

// SNSHttpTestConnectionDetailParam SNSHttpTestConnection detail param
type SNSHttpTestConnectionDetailParam struct {
	Url string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// SNSHttpTestConnectionParam SNSHttpTestConnection request param
type SNSHttpTestConnectionParam struct {
	BaseParam
	Params SNSHttpTestConnectionDetailParam `json:"params"`
}
