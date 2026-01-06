// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CbtTaskInventoryView CbtTask
type CbtTaskInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	ResourceRefs []CbtTaskResourceRefInventoryView `json:"resourceRefs,omitempty"`
}

// CreateCbtTaskEventView CreateCbtTaskEvent
type CreateCbtTaskEventView struct {
	Inventory CbtTaskInventoryView `json:"inventory,omitempty"`
}

// DeleteCbtTaskEventView DeleteCbtTaskEvent
type DeleteCbtTaskEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryCbtTaskView QueryCbtTask
type QueryCbtTaskView struct {
	Inventories []CbtTaskInventoryView `json:"inventories,omitempty"`
}

