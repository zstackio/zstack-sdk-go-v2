// Copyright (c) ZStack.io, Inc.

package param

// DeleteSSORedirectTemplateDetailParam DeleteSSORedirectTemplate detail param
type DeleteSSORedirectTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSSORedirectTemplateParam DeleteSSORedirectTemplate request param
type DeleteSSORedirectTemplateParam struct {
	BaseParam
	Params DeleteSSORedirectTemplateDetailParam `json:"params"`
}
