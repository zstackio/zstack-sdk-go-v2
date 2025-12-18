// Copyright (c) ZStack.io, Inc.

package param

// UpdateSSORedirectTemplateDetailParam UpdateSSORedirectTemplate detail param
type UpdateSSORedirectTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	RedirectTemplate string `json:"redirectTemplate" validate:"required"`
}

// UpdateSSORedirectTemplateParam UpdateSSORedirectTemplate request param
type UpdateSSORedirectTemplateParam struct {
	BaseParam
	Params UpdateSSORedirectTemplateDetailParam `json:"params"`
}
