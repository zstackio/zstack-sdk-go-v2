// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateSAML2ClientParamDetail UpdateSAML2Client detail param
type UpdateSAML2ClientParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	IdpMetadataBase64 string `json:"idpMetadataBase64,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}

// UpdateSAML2ClientParam UpdateSAML2Client request param
type UpdateSAML2ClientParam struct {
	BaseParam
	UpdateSAML2Client UpdateSAML2ClientParamDetail `json:"updateSAML2Client"`
}
// CreateSAML2ClientParamDetail CreateSAML2Client detail param
type CreateSAML2ClientParamDetail struct {
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
	CreateSAML2Client CreateSAML2ClientParamDetail `json:"createSAML2Client"`
}
