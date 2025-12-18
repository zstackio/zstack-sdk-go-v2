// Copyright (c) ZStack.io, Inc.

package param

// UpdateOAuthClientDetailParam UpdateOAuthClient detail param
type UpdateOAuthClientDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`
	TokenUrl string `json:"tokenUrl,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
	UserinfoUrl string `json:"userinfoUrl,omitempty"`
	IdentityProvider string `json:"identityProvider,omitempty"`
	PluginUuid string `json:"pluginUuid,omitempty"`
	LoginType string `json:"loginType,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
	ScopeList []string `json:"scopeList,omitempty"`
}

// UpdateOAuthClientParam UpdateOAuthClient request param
type UpdateOAuthClientParam struct {
	BaseParam
	Params UpdateOAuthClientDetailParam `json:"params"`
}
