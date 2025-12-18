// Copyright (c) ZStack.io, Inc.

package param

// DecodeStackTemplateDetailParam DecodeStackTemplate详细参数
type DecodeStackTemplateDetailParam struct {
	rest string `json:"type,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"parameters,omitempty"`
	rest string `json:"preparameters,omitempty"`
}

// DecodeStackTemplateParam DecodeStackTemplate请求参数
type DecodeStackTemplateParam struct {
	BaseParam
	Params DecodeStackTemplateDetailParam `json:"params"` // 详细参数
}

