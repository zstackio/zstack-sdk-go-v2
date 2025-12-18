// Copyright (c) ZStack.io, Inc.

package param

// ChangeHostStateDetailParam ChangeHostState详细参数
type ChangeHostStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeHostStateParam ChangeHostState请求参数
type ChangeHostStateParam struct {
	BaseParam
	Params ChangeHostStateDetailParam `json:"params"` // 详细参数
}

