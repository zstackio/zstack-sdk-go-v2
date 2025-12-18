// Copyright (c) ZStack.io, Inc.

package param

// SubscribeSNSTopicDetailParam SubscribeSNSTopic详细参数
type SubscribeSNSTopicDetailParam struct {
	rest string `json:"topicUuid" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
}

// SubscribeSNSTopicParam SubscribeSNSTopic请求参数
type SubscribeSNSTopicParam struct {
	BaseParam
	Params SubscribeSNSTopicDetailParam `json:"params"` // 详细参数
}

