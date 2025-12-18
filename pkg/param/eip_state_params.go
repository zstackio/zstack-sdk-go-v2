// Copyright (c) ZStack.io, Inc.

package param

// ChangeEipStateDetailParam ChangeEipState详细参数
type ChangeEipStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeEipStateParam ChangeEipState请求参数
type ChangeEipStateParam struct {
	BaseParam
	Params ChangeEipStateDetailParam `json:"params"` // 详细参数
}

