// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2VirtualIDGroupDetailParam UpdateIAM2VirtualIDGroup详细参数
type UpdateIAM2VirtualIDGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateIAM2VirtualIDGroupParam UpdateIAM2VirtualIDGroup请求参数
type UpdateIAM2VirtualIDGroupParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDGroupDetailParam `json:"params"` // 详细参数
}

