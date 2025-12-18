// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccessKeyStateDetailParam ChangeAccessKeyState详细参数
type ChangeAccessKeyStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeAccessKeyStateParam ChangeAccessKeyState请求参数
type ChangeAccessKeyStateParam struct {
	BaseParam
	Params ChangeAccessKeyStateDetailParam `json:"params"` // 详细参数
}

