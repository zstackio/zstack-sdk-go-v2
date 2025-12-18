// Copyright (c) ZStack.io, Inc.

package param

// PowerResetHostDetailParam PowerResetHost详细参数
type PowerResetHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"returnEarly,omitempty"`
	rest string `json:"method,omitempty"`
}

// PowerResetHostParam PowerResetHost请求参数
type PowerResetHostParam struct {
	BaseParam
	Params PowerResetHostDetailParam `json:"params"` // 详细参数
}

