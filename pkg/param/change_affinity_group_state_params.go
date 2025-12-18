// Copyright (c) ZStack.io, Inc.

package param

// ChangeAffinityGroupStateDetailParam ChangeAffinityGroupState detail param
type ChangeAffinityGroupStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAffinityGroupStateParam ChangeAffinityGroupState request param
type ChangeAffinityGroupStateParam struct {
	BaseParam
	Params ChangeAffinityGroupStateDetailParam `json:"params"`
}
