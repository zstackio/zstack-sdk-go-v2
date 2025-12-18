// Copyright (c) ZStack.io, Inc.

package param

// SNSFeiShuTestConnectionDetailParam SNSFeiShuTestConnection详细参数
type SNSFeiShuTestConnectionDetailParam struct {
	rest string `json:"url,omitempty"`
	rest bool `json:"atAll,omitempty"`
	rest []string `json:"atPersonUserIds,omitempty"`
	rest string `json:"secret,omitempty"`
	rest string `json:"testMsg" validate:"required"` // 必填
	rest string `json:"endpointUuid,omitempty"`
}

// SNSFeiShuTestConnectionParam SNSFeiShuTestConnection请求参数
type SNSFeiShuTestConnectionParam struct {
	BaseParam
	Params SNSFeiShuTestConnectionDetailParam `json:"params"` // 详细参数
}

