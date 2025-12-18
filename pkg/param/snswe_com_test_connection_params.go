// Copyright (c) ZStack.io, Inc.

package param

// SNSWeComTestConnectionDetailParam SNSWeComTestConnection详细参数
type SNSWeComTestConnectionDetailParam struct {
	rest string `json:"url,omitempty"`
	rest bool `json:"atAll,omitempty"`
	rest []string `json:"atPersonUserIds,omitempty"`
	rest string `json:"testMsg" validate:"required"` // 必填
	rest string `json:"endpointUuid,omitempty"`
}

// SNSWeComTestConnectionParam SNSWeComTestConnection请求参数
type SNSWeComTestConnectionParam struct {
	BaseParam
	Params SNSWeComTestConnectionDetailParam `json:"params"` // 详细参数
}

