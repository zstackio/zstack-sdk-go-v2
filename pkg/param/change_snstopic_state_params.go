// Copyright (c) ZStack.io, Inc.

package param

// ChangeSNSTopicStateDetailParam ChangeSNSTopicState detail param
type ChangeSNSTopicStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSNSTopicStateParam ChangeSNSTopicState request param
type ChangeSNSTopicStateParam struct {
	BaseParam
	Params ChangeSNSTopicStateDetailParam `json:"params"`
}
