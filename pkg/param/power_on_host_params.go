// Copyright (c) ZStack.io, Inc.

package param

// PowerOnHostDetailParam PowerOnHost详细参数
type PowerOnHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"returnEarly,omitempty"`
}

// PowerOnHostParam PowerOnHost请求参数
type PowerOnHostParam struct {
	BaseParam
	Params PowerOnHostDetailParam `json:"params"` // 详细参数
}

