// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AgentVersionInventoryView AgentVersion
type AgentVersionInventoryView struct {
	BaseInfoView
	BaseTimeView
	AgentType *string `json:"agentType,omitempty"`
	CurrentVersion *string `json:"currentVersion,omitempty"`
	ExpectVersion *string `json:"expectVersion,omitempty"`
}

// QueryAgentVersionView QueryAgentVersion
type QueryAgentVersionView struct {
	Inventories []AgentVersionInventoryView `json:"inventories,omitempty"`
}

