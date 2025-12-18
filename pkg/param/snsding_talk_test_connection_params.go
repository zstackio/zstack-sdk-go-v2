// Copyright (c) ZStack.io, Inc.

package param

// SNSDingTalkTestConnectionDetailParam SNSDingTalkTestConnection详细参数
type SNSDingTalkTestConnectionDetailParam struct {
	rest string `json:"url,omitempty"`
	rest bool `json:"atAll,omitempty"`
	rest []string `json:"atPersonPhoneNumbers,omitempty"`
	rest string `json:"secret,omitempty"`
	rest string `json:"testMsg" validate:"required"` // 必填
	rest string `json:"endpointUuid,omitempty"`
}

// SNSDingTalkTestConnectionParam SNSDingTalkTestConnection请求参数
type SNSDingTalkTestConnectionParam struct {
	BaseParam
	Params SNSDingTalkTestConnectionDetailParam `json:"params"` // 详细参数
}

