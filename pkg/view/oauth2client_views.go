// Copyright (c) ZStack.io, Inc.

package view

import "time"

// OAuth2ClientInventoryView OAuth2Client
type OAuth2ClientInventoryView struct {
	rest string `json:"clientId,omitempty"`
	rest string `json:"clientSecret,omitempty"`
	rest string `json:"authorizationUrl,omitempty"`
	rest string `json:"tokenUrl,omitempty"`
	rest string `json:"userinfoUrl,omitempty"`
	rest string `json:"grantType,omitempty"`
	rest string `json:"identityProvider,omitempty"`
	rest string `json:"pluginUuid,omitempty"`
	rest string `json:"logoutUrl,omitempty"`
	rest []string `json:"scopeList,omitempty"`
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

