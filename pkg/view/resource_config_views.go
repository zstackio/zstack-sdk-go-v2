// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ResourceConfigInventoryView ResourceConfig
type ResourceConfigInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category string `json:"category,omitempty"`
	Value string `json:"value,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

