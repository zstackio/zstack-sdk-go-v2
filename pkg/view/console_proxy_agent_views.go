// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ConsoleProxyAgentInventoryView ConsoleProxyAgent
type ConsoleProxyAgentInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"consoleProxyOverriddenIp,omitempty"`
	rest int `json:"consoleProxyPort,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

