// Copyright (c) ZStack.io, Inc.

package param

// UnsubscribeEventDetailParam UnsubscribeEvent详细参数
type UnsubscribeEventDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// UnsubscribeEventParam UnsubscribeEvent请求参数
type UnsubscribeEventParam struct {
	BaseParam
	Params UnsubscribeEventDetailParam `json:"params"` // 详细参数
}

