// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateUserProxyConfigParamDetail UpdateUserProxyConfig detail param
type UpdateUserProxyConfigParamDetail struct {
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
	Params UpdateUserProxyConfigParamDetail `json:"params"`
}
// DeleteUserProxyConfigParamDetail DeleteUserProxyConfig detail param
type DeleteUserProxyConfigParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteUserProxyConfigParam DeleteUserProxyConfig request param
type DeleteUserProxyConfigParam struct {
	BaseParam
	Params DeleteUserProxyConfigParamDetail `json:"params"`
}
// CreateUserProxyConfigParamDetail CreateUserProxyConfig detail param
type CreateUserProxyConfigParamDetail struct {
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
	Params CreateUserProxyConfigParamDetail `json:"params"`
}
