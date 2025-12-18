// Copyright (c) ZStack.io, Inc.

package param

// ShutdownHostDetailParam ShutdownHost详细参数
type ShutdownHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"returnEarly,omitempty"`
	rest bool `json:"force,omitempty"`
	rest string `json:"method,omitempty"`
}

// ShutdownHostParam ShutdownHost请求参数
type ShutdownHostParam struct {
	BaseParam
	Params ShutdownHostDetailParam `json:"params"` // 详细参数
}

