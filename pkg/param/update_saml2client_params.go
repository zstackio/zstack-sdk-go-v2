// Copyright (c) ZStack.io, Inc.

package param

// UpdateSAML2ClientDetailParam UpdateSAML2Client detail param
type UpdateSAML2ClientDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	IdpMetadataBase64 string `json:"idpMetadataBase64,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}

// UpdateSAML2ClientParam UpdateSAML2Client request param
type UpdateSAML2ClientParam struct {
	BaseParam
	Params UpdateSAML2ClientDetailParam `json:"params"`
}
