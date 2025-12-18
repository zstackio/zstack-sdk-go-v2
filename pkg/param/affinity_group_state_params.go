// Copyright (c) ZStack.io, Inc.

package param

// ChangeAffinityGroupStateDetailParam ChangeAffinityGroupState详细参数
type ChangeAffinityGroupStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeAffinityGroupStateParam ChangeAffinityGroupState请求参数
type ChangeAffinityGroupStateParam struct {
	BaseParam
	Params ChangeAffinityGroupStateDetailParam `json:"params"` // 详细参数
}

