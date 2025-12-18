// Copyright (c) ZStack.io, Inc.

package param

// ChangeImageStateDetailParam ChangeImageState详细参数
type ChangeImageStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeImageStateParam ChangeImageState请求参数
type ChangeImageStateParam struct {
	BaseParam
	Params ChangeImageStateDetailParam `json:"params"` // 详细参数
}

