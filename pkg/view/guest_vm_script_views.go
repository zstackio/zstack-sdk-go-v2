// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// GuestVmScriptInventoryView GuestVmScript
type GuestVmScriptInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	EncodingType string `json:"encodingType,omitempty"`
	ScriptContent string `json:"scriptContent,omitempty"`
	RenderParams string `json:"renderParams,omitempty"`
	Platform string `json:"platform,omitempty"`
	ScriptType string `json:"scriptType,omitempty"`
	ScriptTimeout int `json:"scriptTimeout,omitempty"`
}

// QueryGuestVmScriptView QueryGuestVmScript
type QueryGuestVmScriptView struct {
	Inventories []GuestVmScriptInventoryView `json:"inventories,omitempty"`
}

// DeleteGuestVmScriptEventView DeleteGuestVmScriptEvent
type DeleteGuestVmScriptEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateGuestVmScriptEventView CreateGuestVmScriptEvent
type CreateGuestVmScriptEventView struct {
	Inventory GuestVmScriptInventoryView `json:"inventory,omitempty"`
}

// UpdateGuestVmScriptEventView UpdateGuestVmScriptEvent
type UpdateGuestVmScriptEventView struct {
	Inventory GuestVmScriptInventoryView `json:"inventory,omitempty"`
}

