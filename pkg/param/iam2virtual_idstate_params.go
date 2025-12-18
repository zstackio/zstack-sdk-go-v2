// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2VirtualIDStateDetailParam ChangeIAM2VirtualIDState详细参数
type ChangeIAM2VirtualIDStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeIAM2VirtualIDStateParam ChangeIAM2VirtualIDState请求参数
type ChangeIAM2VirtualIDStateParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDStateDetailParam `json:"params"` // 详细参数
}

