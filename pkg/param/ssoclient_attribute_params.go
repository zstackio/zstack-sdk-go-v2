// Copyright (c) ZStack.io, Inc.

package param

// UpdateSSOClientAttributeDetailParam UpdateSSOClientAttribute详细参数
type UpdateSSOClientAttributeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"value,omitempty"`
	rest string `json:"purpose,omitempty"`
	rest string `json:"type,omitempty"`
}

// UpdateSSOClientAttributeParam UpdateSSOClientAttribute请求参数
type UpdateSSOClientAttributeParam struct {
	BaseParam
	Params UpdateSSOClientAttributeDetailParam `json:"params"` // 详细参数
}

