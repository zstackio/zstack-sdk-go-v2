// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ThirdpartyPlatformInventoryView ThirdpartyPlatform
type ThirdpartyPlatformInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Template string `json:"template,omitempty"`
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	LastSyncDate ZStackTime `json:"lastSyncDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
}

// QueryThirdpartyPlatformView QueryThirdpartyPlatform
type QueryThirdpartyPlatformView struct {
	Inventories []ThirdpartyPlatformInventoryView `json:"inventories,omitempty"`
}

// UpdateThirdpartyPlatformEventView UpdateThirdpartyPlatformEvent
type UpdateThirdpartyPlatformEventView struct {
	Inventory ThirdpartyPlatformInventoryView `json:"inventory,omitempty"`
}

// AddThirdpartyPlatformEventView AddThirdpartyPlatformEvent
type AddThirdpartyPlatformEventView struct {
	Inventory ThirdpartyPlatformInventoryView `json:"inventory,omitempty"`
}

// DeleteThirdpartyPlatformEventView DeleteThirdpartyPlatformEvent
type DeleteThirdpartyPlatformEventView struct {
	Success bool `json:"success,omitempty"`
}

