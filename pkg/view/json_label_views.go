// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// JsonLabelInventoryView JsonLabel
type JsonLabelInventoryView struct {
	Id int64 `json:"id,omitempty"`
	LabelKey *string `json:"labelKey,omitempty"`
	LabelValue *string `json:"labelValue,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateLogConfigurationEventView UpdateLogConfigurationEvent
type UpdateLogConfigurationEventView struct {
	Inventory JsonLabelInventoryView `json:"inventory,omitempty"`
}

// GetLogConfigurationView GetLogConfiguration
type GetLogConfigurationView struct {
	Inventories []JsonLabelInventoryView `json:"inventories,omitempty"`
}

// AddLogConfigurationEventView AddLogConfigurationEvent
type AddLogConfigurationEventView struct {
	Inventory JsonLabelInventoryView `json:"inventory,omitempty"`
}

