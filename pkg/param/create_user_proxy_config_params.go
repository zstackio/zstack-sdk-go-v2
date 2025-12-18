// Copyright (c) ZStack.io, Inc.

package param

// CreateUserProxyConfigDetailParam CreateUserProxyConfig detail param
type CreateUserProxyConfigDetailParam struct {
	ProxyType string `json:"proxyType" validate:"required"`
	ProxyHost string `json:"proxyHost" validate:"required"`
	ProxyPort int `json:"proxyPort" validate:"required"`
	ProxyUsername string `json:"proxyUsername,omitempty"`
	ProxyPassword string `json:"proxyPassword,omitempty"`
	IsEnabled bool `json:"isEnabled,omitempty"`
	ProxyProtocolVersion string `json:"proxyProtocolVersion,omitempty"`
	UseSsl bool `json:"useSsl,omitempty"`
	NoProxy string `json:"noProxy,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateUserProxyConfigParam CreateUserProxyConfig request param
type CreateUserProxyConfigParam struct {
	BaseParam
	Params CreateUserProxyConfigDetailParam `json:"params"`
}
