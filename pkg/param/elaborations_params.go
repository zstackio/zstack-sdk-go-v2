// Copyright (c) ZStack.io, Inc.

package param

// GetElaborationsDetailParam GetElaborations详细参数
type GetElaborationsDetailParam struct {
	rest string `json:"category,omitempty"`
	rest string `json:"code,omitempty"`
	rest string `json:"regex,omitempty"`
}

// GetElaborationsParam GetElaborations请求参数
type GetElaborationsParam struct {
	BaseParam
	Params GetElaborationsDetailParam `json:"params"` // 详细参数
}

