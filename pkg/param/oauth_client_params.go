// Copyright (c) ZStack.io, Inc.

package param

// UpdateOAuthClientDetailParam UpdateOAuthClient详细参数
type UpdateOAuthClientDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"clientId,omitempty"`
	rest string `json:"clientSecret,omitempty"`
	rest string `json:"authorizationUrl,omitempty"`
	rest string `json:"tokenUrl,omitempty"`
	rest string `json:"redirectUrl,omitempty"`
	rest string `json:"userinfoUrl,omitempty"`
	rest string `json:"identityProvider,omitempty"`
	rest string `json:"pluginUuid,omitempty"`
	rest string `json:"loginType,omitempty"`
	rest string `json:"logoutUrl,omitempty"`
	rest []string `json:"scopeList,omitempty"`
}

// UpdateOAuthClientParam UpdateOAuthClient请求参数
type UpdateOAuthClientParam struct {
	BaseParam
	Params UpdateOAuthClientDetailParam `json:"params"` // 详细参数
}

