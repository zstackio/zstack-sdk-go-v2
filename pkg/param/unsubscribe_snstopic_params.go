// Copyright (c) ZStack.io, Inc.

package param

// UnsubscribeSNSTopicDetailParam UnsubscribeSNSTopic详细参数
type UnsubscribeSNSTopicDetailParam struct {
	rest string `json:"topicUuid" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
}

// UnsubscribeSNSTopicParam UnsubscribeSNSTopic请求参数
type UnsubscribeSNSTopicParam struct {
	BaseParam
	Params UnsubscribeSNSTopicDetailParam `json:"params"` // 详细参数
}

