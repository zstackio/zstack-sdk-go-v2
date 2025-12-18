// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PluginDriverInventoryView PluginDriver
type PluginDriverInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"features,omitempty"`
	rest []interface{} `json:"optionTypes,omitempty"`
	rest bool `json:"deleted,omitempty"`
	rest string `json:"license,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

