// Copyright (c) ZStack.io, Inc.

package view

// ExportImageFromBackupStorageEventView ExportImageFromBackupStorageEvent
type ExportImageFromBackupStorageEventView struct {
	ImageUrl string `json:"imageUrl,omitempty"`
	ExportMd5Sum string `json:"exportMd5Sum,omitempty"`
	Success bool `json:"success,omitempty"`
}

