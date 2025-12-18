// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AgentVersionInventoryView AgentVersion
type AgentVersionInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"agentType,omitempty"`
	rest string `json:"currentVersion,omitempty"`
	rest string `json:"expectVersion,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

