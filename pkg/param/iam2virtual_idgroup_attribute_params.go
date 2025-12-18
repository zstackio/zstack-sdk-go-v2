// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2VirtualIDGroupAttributeDetailParam UpdateIAM2VirtualIDGroupAttribute详细参数
type UpdateIAM2VirtualIDGroupAttributeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"value" validate:"required"` // 必填
}

// UpdateIAM2VirtualIDGroupAttributeParam UpdateIAM2VirtualIDGroupAttribute请求参数
type UpdateIAM2VirtualIDGroupAttributeParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDGroupAttributeDetailParam `json:"params"` // 详细参数
}

