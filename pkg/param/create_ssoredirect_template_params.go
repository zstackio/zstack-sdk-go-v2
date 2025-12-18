// Copyright (c) ZStack.io, Inc.

package param

// CreateSSORedirectTemplateDetailParam CreateSSORedirectTemplate detail param
type CreateSSORedirectTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ClientUuid string `json:"clientUuid" validate:"required"`
	RedirectTemplate string `json:"redirectTemplate" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSSORedirectTemplateParam CreateSSORedirectTemplate request param
type CreateSSORedirectTemplateParam struct {
	BaseParam
	Params CreateSSORedirectTemplateDetailParam `json:"params"`
}
