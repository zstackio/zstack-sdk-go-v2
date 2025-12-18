// Copyright (c) ZStack.io, Inc.

package view

import "time"

// GuestVmScriptInventoryView GuestVmScript
type GuestVmScriptInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"encodingType,omitempty"`
	rest string `json:"scriptContent,omitempty"`
	rest string `json:"renderParams,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"scriptType,omitempty"`
	rest int `json:"scriptTimeout,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

