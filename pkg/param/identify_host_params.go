// Copyright (c) ZStack.io, Inc.

package param

// IdentifyHostDetailParam IdentifyHost详细参数
type IdentifyHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int64 `json:"interval,omitempty"`
}

// IdentifyHostParam IdentifyHost请求参数
type IdentifyHostParam struct {
	BaseParam
	Params IdentifyHostDetailParam `json:"params"` // 详细参数
}

