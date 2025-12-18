// Copyright (c) ZStack.io, Inc.

package view

// GetResourceConfigView GetResourceConfig
type GetResourceConfigView struct {
	Value string `json:"value,omitempty"`
	EffectiveConfigs []ResourceConfigInventoryView `json:"effectiveConfigs,omitempty"`
	Success bool `json:"success,omitempty"`
}

