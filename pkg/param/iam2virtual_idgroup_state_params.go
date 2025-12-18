// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2VirtualIDGroupStateDetailParam ChangeIAM2VirtualIDGroupState详细参数
type ChangeIAM2VirtualIDGroupStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeIAM2VirtualIDGroupStateParam ChangeIAM2VirtualIDGroupState请求参数
type ChangeIAM2VirtualIDGroupStateParam struct {
	BaseParam
	Params ChangeIAM2VirtualIDGroupStateDetailParam `json:"params"` // 详细参数
}

