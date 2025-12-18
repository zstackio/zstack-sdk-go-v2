// Copyright (c) ZStack.io, Inc.

package param

// ChangePreconfigurationTemplateStateDetailParam ChangePreconfigurationTemplateState详细参数
type ChangePreconfigurationTemplateStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangePreconfigurationTemplateStateParam ChangePreconfigurationTemplateState请求参数
type ChangePreconfigurationTemplateStateParam struct {
	BaseParam
	Params ChangePreconfigurationTemplateStateDetailParam `json:"params"` // 详细参数
}

