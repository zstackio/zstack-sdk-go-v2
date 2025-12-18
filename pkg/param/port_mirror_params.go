// Copyright (c) ZStack.io, Inc.

package param

// DeletePortMirrorDetailParam DeletePortMirror详细参数
type DeletePortMirrorDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeletePortMirrorParam DeletePortMirror请求参数
type DeletePortMirrorParam struct {
	BaseParam
	Params DeletePortMirrorDetailParam `json:"params"` // 详细参数
}

