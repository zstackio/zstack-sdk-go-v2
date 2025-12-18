// Copyright (c) ZStack.io, Inc.

package param

// ChangeSNSTopicStateDetailParam ChangeSNSTopicState详细参数
type ChangeSNSTopicStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeSNSTopicStateParam ChangeSNSTopicState请求参数
type ChangeSNSTopicStateParam struct {
	BaseParam
	Params ChangeSNSTopicStateDetailParam `json:"params"` // 详细参数
}

