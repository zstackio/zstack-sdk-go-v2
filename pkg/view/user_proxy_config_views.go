// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserProxyConfigInventoryView UserProxyConfig
type UserProxyConfigInventoryView struct {
	Uuid                 string    `json:"uuid,omitempty"`
	ProxyType            string    `json:"proxyType,omitempty"`
	ProxyHost            string    `json:"proxyHost,omitempty"`
	ProxyPort            int       `json:"proxyPort,omitempty"`
	ProxyUsername        string    `json:"proxyUsername,omitempty"`
	ProxyPassword        string    `json:"proxyPassword,omitempty"`
	IsEnabled            bool      `json:"isEnabled,omitempty"`
	ProxyProtocolVersion string    `json:"proxyProtocolVersion,omitempty"`
	UseSsl               bool      `json:"useSsl,omitempty"`
	NoProxy              string    `json:"noProxy,omitempty"`
	CreateDate           time.Time `json:"createDate,omitempty"`
	LastOpDate           time.Time `json:"lastOpDate,omitempty"`
}

// UpdateUserProxyConfigEventView UpdateUserProxyConfigEvent
type UpdateUserProxyConfigEventView struct {
	Inventory UserProxyConfigInventoryView `json:"inventory,omitempty"`
}

// DeleteUserProxyConfigEventView DeleteUserProxyConfigEvent
type DeleteUserProxyConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryUserProxyConfigView QueryUserProxyConfig
type QueryUserProxyConfigView struct {
	Inventories []UserProxyConfigInventoryView `json:"inventories,omitempty"`
}

// CreateUserProxyConfigEventView CreateUserProxyConfigEvent
type CreateUserProxyConfigEventView struct {
	Inventory UserProxyConfigInventoryView `json:"inventory,omitempty"`
}
