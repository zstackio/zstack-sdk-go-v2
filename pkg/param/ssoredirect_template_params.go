// Copyright (c) ZStack.io, Inc.

package param

// DeleteSSORedirectTemplateDetailParam DeleteSSORedirectTemplate详细参数
type DeleteSSORedirectTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteSSORedirectTemplateParam DeleteSSORedirectTemplate请求参数
type DeleteSSORedirectTemplateParam struct {
	BaseParam
	Params DeleteSSORedirectTemplateDetailParam `json:"params"` // 详细参数
}

