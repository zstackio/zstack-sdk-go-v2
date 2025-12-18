// Copyright (c) ZStack.io, Inc.

package param

// ChangeL3NetworkStateDetailParam ChangeL3NetworkState详细参数
type ChangeL3NetworkStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeL3NetworkStateParam ChangeL3NetworkState请求参数
type ChangeL3NetworkStateParam struct {
	BaseParam
	Params ChangeL3NetworkStateDetailParam `json:"params"` // 详细参数
}

