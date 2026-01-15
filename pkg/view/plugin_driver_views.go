// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PluginDriverInventoryView PluginDriver
type PluginDriverInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Features string `json:"features,omitempty"`
	OptionTypes []OptionTypeView `json:"optionTypes,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	License string `json:"license,omitempty"`
	Version string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
}

// QueryPluginDriversView QueryPluginDrivers
type QueryPluginDriversView struct {
	Inventories []PluginDriverInventoryView `json:"inventories,omitempty"`
}

