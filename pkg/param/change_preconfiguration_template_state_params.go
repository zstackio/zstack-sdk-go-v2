// Copyright (c) ZStack.io, Inc.

package param

// ChangePreconfigurationTemplateStateDetailParam ChangePreconfigurationTemplateState detail param
type ChangePreconfigurationTemplateStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePreconfigurationTemplateStateParam ChangePreconfigurationTemplateState request param
type ChangePreconfigurationTemplateStateParam struct {
	BaseParam
	Params ChangePreconfigurationTemplateStateDetailParam `json:"params"`
}
