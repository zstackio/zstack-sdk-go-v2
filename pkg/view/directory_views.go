// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DirectoryInventoryView Directory
type DirectoryInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	RootDirectoryUuid string `json:"rootDirectoryUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// UpdateDirectoryEventView UpdateDirectoryEvent
type UpdateDirectoryEventView struct {
	Inventory DirectoryInventoryView `json:"inventory,omitempty"`
}

// CreateDirectoryEventView CreateDirectoryEvent
type CreateDirectoryEventView struct {
	Inventory DirectoryInventoryView `json:"inventory,omitempty"`
}

// QueryDirectoryView QueryDirectory
type QueryDirectoryView struct {
	Inventories []DirectoryInventoryView `json:"inventories,omitempty"`
}

// DeleteDirectoryEventView DeleteDirectoryEvent
type DeleteDirectoryEventView struct {
	Success bool `json:"success,omitempty"`
}

