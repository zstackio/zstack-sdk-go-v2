// Copyright (c) ZStack.io, Inc.

package param

// CreateSAML2ClientDetailParam CreateSAML2Client detail param
type CreateSAML2ClientDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	IdpMetadataBase64 string `json:"idpMetadataBase64" validate:"required"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
	LoginType string `json:"loginType" validate:"required"`
	UrlTemplate string `json:"urlTemplate" validate:"required"`
	Attributes []ExtendedAttributeParam `json:"attributes,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSAML2ClientParam CreateSAML2Client request param
type CreateSAML2ClientParam struct {
	BaseParam
	Params CreateSAML2ClientDetailParam `json:"params"`
}
