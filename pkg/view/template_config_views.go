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

