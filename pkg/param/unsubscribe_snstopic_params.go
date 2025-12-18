// Copyright (c) ZStack.io, Inc.

package param

// UnsubscribeSNSTopicDetailParam UnsubscribeSNSTopic detail param
type UnsubscribeSNSTopicDetailParam struct {
	TopicUuid string `json:"topicUuid" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
}

// UnsubscribeSNSTopicParam UnsubscribeSNSTopic request param
type UnsubscribeSNSTopicParam struct {
	BaseParam
	Params UnsubscribeSNSTopicDetailParam `json:"params"`
}
