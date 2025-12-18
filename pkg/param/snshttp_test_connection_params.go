// Copyright (c) ZStack.io, Inc.

package param

// SNSHttpTestConnectionDetailParam SNSHttpTestConnection详细参数
type SNSHttpTestConnectionDetailParam struct {
	rest string `json:"url,omitempty"`
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
}

// SNSHttpTestConnectionParam SNSHttpTestConnection请求参数
type SNSHttpTestConnectionParam struct {
	BaseParam
	Params SNSHttpTestConnectionDetailParam `json:"params"` // 详细参数
}

