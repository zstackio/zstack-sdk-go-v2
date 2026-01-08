// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GlobalConfigTemplateInventoryView GlobalConfigTemplate
type GlobalConfigTemplateInventoryView struct {
	Uuid        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// QueryGlobalConfigTemplateView QueryGlobalConfigTemplate
type QueryGlobalConfigTemplateView struct {
	Inventories []GlobalConfigTemplateInventoryView `json:"inventories,omitempty"`
}
