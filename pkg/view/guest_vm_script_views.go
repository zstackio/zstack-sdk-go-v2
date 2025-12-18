// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GuestVmScriptInventoryView GuestVmScript
type GuestVmScriptInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	EncodingType string `json:"encodingType,omitempty"`
	ScriptContent string `json:"scriptContent,omitempty"`
	RenderParams string `json:"renderParams,omitempty"`
	Platform string `json:"platform,omitempty"`
	ScriptType string `json:"scriptType,omitempty"`
	ScriptTimeout int `json:"scriptTimeout,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

