// Copyright (c) ZStack.io, Inc.

package param

// SNSEmailTestConnectionDetailParam SNSEmailTestConnection详细参数
type SNSEmailTestConnectionDetailParam struct {
	rest []string `json:"emails,omitempty"`
	rest string `json:"platformUuid,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest string `json:"subject,omitempty"`
	rest string `json:"text,omitempty"`
}

// SNSEmailTestConnectionParam SNSEmailTestConnection请求参数
type SNSEmailTestConnectionParam struct {
	BaseParam
	Params SNSEmailTestConnectionDetailParam `json:"params"` // 详细参数
}

