// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2ProjectStateDetailParam ChangeIAM2ProjectState详细参数
type ChangeIAM2ProjectStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeIAM2ProjectStateParam ChangeIAM2ProjectState请求参数
type ChangeIAM2ProjectStateParam struct {
	BaseParam
	Params ChangeIAM2ProjectStateDetailParam `json:"params"` // 详细参数
}

