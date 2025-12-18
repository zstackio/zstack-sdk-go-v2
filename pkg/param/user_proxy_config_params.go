// Copyright (c) ZStack.io, Inc.

package param

// UpdateUserProxyConfigDetailParam UpdateUserProxyConfig详细参数
type UpdateUserProxyConfigDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"proxyType,omitempty"`
	rest string `json:"proxyHost,omitempty"`
	rest int `json:"proxyPort,omitempty"`
	rest string `json:"proxyUsername,omitempty"`
	rest string `json:"proxyPassword,omitempty"`
	rest bool `json:"isEnabled,omitempty"`
	rest string `json:"proxyProtocolVersion,omitempty"`
	rest bool `json:"useSsl,omitempty"`
	rest string `json:"noProxy,omitempty"`
}

// UpdateUserProxyConfigParam UpdateUserProxyConfig请求参数
type UpdateUserProxyConfigParam struct {
	BaseParam
	Params UpdateUserProxyConfigDetailParam `json:"params"` // 详细参数
}

