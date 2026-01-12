// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OAuth2ClientInventoryView OAuth2Client
type OAuth2ClientInventoryView struct {
	BaseInfoView
	BaseTimeView
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
	TokenUrl *string `json:"tokenUrl,omitempty"`
	UserinfoUrl *string `json:"userinfoUrl,omitempty"`
	GrantType *string `json:"grantType,omitempty"`
	IdentityProvider *string `json:"identityProvider,omitempty"`
	PluginUuid *string `json:"pluginUuid,omitempty"`
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	ScopeList []string `json:"scopeList,omitempty"`
	Description *string `json:"description,omitempty"`
	ClientType *string `json:"clientType,omitempty"`
	LoginType *string `json:"loginType,omitempty"`
	LoginMNUrl *string `json:"loginMNUrl,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	Attributes []SSOClientAttributeInventoryView `json:"attributes,omitempty"`
}

// CreateOAuthClientEventView CreateOAuthClientEvent
type CreateOAuthClientEventView struct {
	Inventory OAuth2ClientInventoryView `json:"inventory,omitempty"`
}

// UpdateOAuthClientEventView UpdateOAuthClientEvent
type UpdateOAuthClientEventView struct {
	Inventory OAuth2ClientInventoryView `json:"inventory,omitempty"`
}

