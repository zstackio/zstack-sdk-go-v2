// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GuestToolsInventoryView GuestTools
type GuestToolsInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	Version string `json:"version,omitempty"`
	AgentType string `json:"agentType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

