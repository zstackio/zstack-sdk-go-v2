// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LogConfigurationInventoryView LogConfiguration
type LogConfigurationInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Type string `json:"type,omitempty"`
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty"`
}

// UpdateLogConfigurationEventView UpdateLogConfigurationEvent
type UpdateLogConfigurationEventView struct {
	Inventory JsonLabelInventoryView `json:"inventory,omitempty"`
}

// GetLogConfigurationView GetLogConfiguration
type GetLogConfigurationView struct {
	Inventories []JsonLabelInventoryView `json:"inventories,omitempty"`
}

// DeleteLogConfigurationEventView DeleteLogConfigurationEvent
type DeleteLogConfigurationEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddLogConfigurationEventView AddLogConfigurationEvent
type AddLogConfigurationEventView struct {
	Inventory JsonLabelInventoryView `json:"inventory,omitempty"`
}

