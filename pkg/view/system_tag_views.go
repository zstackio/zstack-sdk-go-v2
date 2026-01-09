// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SystemTagInventoryView SystemTag
type SystemTagInventoryView struct {
	Inherent *bool `json:"inherent,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Type *string `json:"type,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateSystemTagEventView UpdateSystemTagEvent
type UpdateSystemTagEventView struct {
	Inventory SystemTagInventoryView `json:"inventory,omitempty"`
}

// CreateSystemTagEventView CreateSystemTagEvent
type CreateSystemTagEventView struct {
	Inventory SystemTagInventoryView `json:"inventory,omitempty"`
}

// QuerySystemTagView QuerySystemTag
type QuerySystemTagView struct {
	Inventories []SystemTagInventoryView `json:"inventories,omitempty"`
}

// CreateSystemTagsEventView CreateSystemTagsEvent
type CreateSystemTagsEventView struct {
	Inventories []SystemTagInventoryView `json:"inventories,omitempty"`
}

