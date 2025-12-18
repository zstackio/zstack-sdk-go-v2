// Copyright (c) ZStack.io, Inc.

package param

// UpdateSAML2ClientDetailParam UpdateSAML2Client详细参数
type UpdateSAML2ClientDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"idpMetadataBase64,omitempty"`
	rest string `json:"redirectUrl,omitempty"`
}

// UpdateSAML2ClientParam UpdateSAML2Client请求参数
type UpdateSAML2ClientParam struct {
	BaseParam
	Params UpdateSAML2ClientDetailParam `json:"params"` // 详细参数
}

