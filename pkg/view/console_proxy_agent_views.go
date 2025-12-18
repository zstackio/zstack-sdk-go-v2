// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ConsoleProxyAgentInventoryView ConsoleProxyAgent
type ConsoleProxyAgentInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	ConsoleProxyOverriddenIp string `json:"consoleProxyOverriddenIp,omitempty"`
	ConsoleProxyPort int `json:"consoleProxyPort,omitempty"`
	Type string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

