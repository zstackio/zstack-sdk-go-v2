// Copyright (c) ZStack.io, Inc.

package param

// UpdateUserProxyConfigDetailParam UpdateUserProxyConfig detail param
type UpdateUserProxyConfigDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ProxyType string `json:"proxyType,omitempty"`
	ProxyHost string `json:"proxyHost,omitempty"`
	ProxyPort int `json:"proxyPort,omitempty"`
	ProxyUsername string `json:"proxyUsername,omitempty"`
	ProxyPassword string `json:"proxyPassword,omitempty"`
	IsEnabled bool `json:"isEnabled,omitempty"`
	ProxyProtocolVersion string `json:"proxyProtocolVersion,omitempty"`
	UseSsl bool `json:"useSsl,omitempty"`
	NoProxy string `json:"noProxy,omitempty"`
}

// UpdateUserProxyConfigParam UpdateUserProxyConfig request param
type UpdateUserProxyConfigParam struct {
	BaseParam
	Params UpdateUserProxyConfigDetailParam `json:"params"`
}
