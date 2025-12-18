// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UserProxyConfigInventoryView UserProxyConfig
type UserProxyConfigInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"proxyType,omitempty"`
	rest string `json:"proxyHost,omitempty"`
	rest int `json:"proxyPort,omitempty"`
	rest string `json:"proxyUsername,omitempty"`
	rest string `json:"proxyPassword,omitempty"`
	rest bool `json:"isEnabled,omitempty"`
	rest string `json:"proxyProtocolVersion,omitempty"`
	rest bool `json:"useSsl,omitempty"`
	rest string `json:"noProxy,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

