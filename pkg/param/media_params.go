// Copyright (c) ZStack.io, Inc.

package param

// DeleteMediaDetailParam DeleteMedia详细参数
type DeleteMediaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteMediaParam DeleteMedia请求参数
type DeleteMediaParam struct {
	BaseParam
	Params DeleteMediaDetailParam `json:"params"` // 详细参数
}

