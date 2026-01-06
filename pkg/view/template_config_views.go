// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TemplateConfigInventoryView TemplateConfig
type TemplateConfigInventoryView struct {
	TemplateUuid string `json:"templateUuid,omitempty"`
	Category string `json:"category,omitempty"`
	Name string `json:"name,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Value string `json:"value,omitempty"`
}

// UpdateTemplateConfigEventView UpdateTemplateConfigEvent
type UpdateTemplateConfigEventView struct {
	Inventory TemplateConfigInventoryView `json:"inventory,omitempty"`
}

// QueryTemplateConfigView QueryTemplateConfig
type QueryTemplateConfigView struct {
	Inventories []TemplateConfigInventoryView `json:"inventories,omitempty"`
}

// RevertTemplateConfigEventView RevertTemplateConfigEvent
type RevertTemplateConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// ApplyTemplateConfigEventView ApplyTemplateConfigEvent
type ApplyTemplateConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

// ResetTemplateConfigEventView ResetTemplateConfigEvent
type ResetTemplateConfigEventView struct {
	Success bool `json:"success,omitempty"`
}

