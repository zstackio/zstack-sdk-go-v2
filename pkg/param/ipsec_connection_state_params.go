// Copyright (c) ZStack.io, Inc.

package param

// ChangeIPSecConnectionStateDetailParam ChangeIPSecConnectionState详细参数
type ChangeIPSecConnectionStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeIPSecConnectionStateParam ChangeIPSecConnectionState请求参数
type ChangeIPSecConnectionStateParam struct {
	BaseParam
	Params ChangeIPSecConnectionStateDetailParam `json:"params"` // 详细参数
}

