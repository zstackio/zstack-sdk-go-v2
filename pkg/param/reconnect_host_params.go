// Copyright (c) ZStack.io, Inc.

package param

// ReconnectHostDetailParam ReconnectHost详细参数
type ReconnectHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectHostParam ReconnectHost请求参数
type ReconnectHostParam struct {
	BaseParam
	Params ReconnectHostDetailParam `json:"params"` // 详细参数
}

