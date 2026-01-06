// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NasFileSystemInventoryView NasFileSystem
type NasFileSystemInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FileSystemId string `json:"fileSystemId,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryNasFileSystemView QueryNasFileSystem
type QueryNasFileSystemView struct {
	Inventories []NasFileSystemInventoryView `json:"inventories,omitempty"`
}

// UpdateNasFileSystemEventView UpdateNasFileSystemEvent
type UpdateNasFileSystemEventView struct {
	Inventory NasFileSystemInventoryView `json:"inventory,omitempty"`
}

// DeleteNasFileSystemEventView DeleteNasFileSystemEvent
type DeleteNasFileSystemEventView struct {
	Success bool `json:"success,omitempty"`
}

