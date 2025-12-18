// Copyright (c) ZStack.io, Inc.

package param

// ChangePortMirrorStateDetailParam ChangePortMirrorState详细参数
type ChangePortMirrorStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangePortMirrorStateParam ChangePortMirrorState请求参数
type ChangePortMirrorStateParam struct {
	BaseParam
	Params ChangePortMirrorStateDetailParam `json:"params"` // 详细参数
}

