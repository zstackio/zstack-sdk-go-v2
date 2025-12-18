// Copyright (c) ZStack.io, Inc.

package param

// SubscribeSNSTopicDetailParam SubscribeSNSTopic detail param
type SubscribeSNSTopicDetailParam struct {
	TopicUuid string `json:"topicUuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
}

// SubscribeSNSTopicParam SubscribeSNSTopic request param
type SubscribeSNSTopicParam struct {
	BaseParam
	Params SubscribeSNSTopicDetailParam `json:"params"`
}
