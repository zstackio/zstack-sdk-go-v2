// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MediaInventoryView Media
type MediaInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
}

// ChangeMediaStateEventView ChangeMediaStateEvent
type ChangeMediaStateEventView struct {
	Inventory MediaInventoryView `json:"inventory,omitempty"`
}

// DeleteMediaEventView DeleteMediaEvent
type DeleteMediaEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryMediaView QueryMedia
type QueryMediaView struct {
	Inventories []MediaInventoryView `json:"inventories,omitempty"`
}

// CreateMediaEventView CreateMediaEvent
type CreateMediaEventView struct {
	Inventory MediaInventoryView `json:"inventory,omitempty"`
}

