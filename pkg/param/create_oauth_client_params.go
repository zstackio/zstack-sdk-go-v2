// Copyright (c) ZStack.io, Inc.

package param

// CreateOAuthClientDetailParam CreateOAuthClient detail param
type CreateOAuthClientDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ClientId string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret,omitempty"`
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`
	TokenUrl string `json:"tokenUrl" validate:"required"`
	UserinfoUrl string `json:"userinfoUrl,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
	LoginType string `json:"loginType" validate:"required"`
	GrantType string `json:"grantType" validate:"required"`
	IdentityProvider string `json:"identityProvider,omitempty"`
	PluginUuid string `json:"pluginUuid,omitempty"`
	UrlTemplate string `json:"urlTemplate" validate:"required"`
	ClientType string `json:"clientType" validate:"required"`
	ScopeList []string `json:"scopeList,omitempty"`
	Attributes []interface{} `json:"attributes,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateOAuthClientParam CreateOAuthClient request param
type CreateOAuthClientParam struct {
	BaseParam
	Params CreateOAuthClientDetailParam `json:"params"`
}
