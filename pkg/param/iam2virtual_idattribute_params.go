// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2VirtualIDAttributeDetailParam UpdateIAM2VirtualIDAttribute详细参数
type UpdateIAM2VirtualIDAttributeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"value" validate:"required"` // 必填
}

// UpdateIAM2VirtualIDAttributeParam UpdateIAM2VirtualIDAttribute请求参数
type UpdateIAM2VirtualIDAttributeParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDAttributeDetailParam `json:"params"` // 详细参数
}

