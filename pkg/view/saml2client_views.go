// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SAML2ClientInventoryView SAML2Client
type SAML2ClientInventoryView struct {
	IdpMetadataBase64 *string `json:"idpMetadataBase64,omitempty"`
	SpX509Certificate *string `json:"spX509Certificate,omitempty"`
	SpMetadataUrl *string `json:"spMetadataUrl,omitempty"`
	State string `json:"state,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ClientType *string `json:"clientType,omitempty"`
	LoginType *string `json:"loginType,omitempty"`
	LoginMNUrl *string `json:"loginMNUrl,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	Attributes []SSOClientAttributeInventoryView `json:"attributes,omitempty"`
}

// UpdateSAML2ClientEventView UpdateSAML2ClientEvent
type UpdateSAML2ClientEventView struct {
	Inventory SAML2ClientInventoryView `json:"inventory,omitempty"`
}

// CreateSAML2ClientEventView CreateSAML2ClientEvent
type CreateSAML2ClientEventView struct {
	Inventory SAML2ClientInventoryView `json:"inventory,omitempty"`
}

