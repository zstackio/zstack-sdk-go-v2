// Copyright (c) ZStack.io, Inc.

package view

import "time"

// GuestToolsInventoryView GuestTools
type GuestToolsInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementNodeUuid,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"agentType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

