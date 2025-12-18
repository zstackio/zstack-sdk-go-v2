// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ConsoleProxyInventoryView ConsoleProxy
type ConsoleProxyInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"agentIp,omitempty"`
	rest string `json:"token,omitempty"`
	rest string `json:"agentType,omitempty"`
	rest string `json:"proxyHostname,omitempty"`
	rest int `json:"proxyPort,omitempty"`
	rest string `json:"targetSchema,omitempty"`
	rest string `json:"targetHostname,omitempty"`
	rest int `json:"targetPort,omitempty"`
	rest string `json:"scheme,omitempty"`
	rest string `json:"proxyIdentity,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"version,omitempty"`
	rest time.Time `json:"expiredDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

