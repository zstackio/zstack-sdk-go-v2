// Copyright (c) ZStack.io, Inc.

package param

// ChangeMediaStateDetailParam ChangeMediaState详细参数
type ChangeMediaStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeMediaStateParam ChangeMediaState请求参数
type ChangeMediaStateParam struct {
	BaseParam
	Params ChangeMediaStateDetailParam `json:"params"` // 详细参数
}

