// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AgentVersionInventoryView AgentVersion
type AgentVersionInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AgentType string `json:"agentType,omitempty"`
	CurrentVersion string `json:"currentVersion,omitempty"`
	ExpectVersion string `json:"expectVersion,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryAgentVersionView QueryAgentVersion
type QueryAgentVersionView struct {
	Inventories []AgentVersionInventoryView `json:"inventories,omitempty"`
}

