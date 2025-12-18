// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SAML2ClientInventoryView SAML2Client
type SAML2ClientInventoryView struct {
	rest string `json:"idpMetadataBase64,omitempty"`
	rest string `json:"spX509Certificate,omitempty"`
	rest string `json:"spMetadataUrl,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"clientType,omitempty"`
	rest string `json:"loginType,omitempty"`
	rest string `json:"loginMNUrl,omitempty"`
	rest string `json:"redirectUrl,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest []SSOClientAttributeInventoryView `json:"attributes,omitempty"`
}

