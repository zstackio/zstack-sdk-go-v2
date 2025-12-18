// Copyright (c) ZStack.io, Inc.

package param

// ChangeSchedulerStateDetailParam ChangeSchedulerState详细参数
type ChangeSchedulerStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeSchedulerStateParam ChangeSchedulerState请求参数
type ChangeSchedulerStateParam struct {
	BaseParam
	Params ChangeSchedulerStateDetailParam `json:"params"` // 详细参数
}

