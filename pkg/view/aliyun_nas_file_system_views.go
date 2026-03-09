// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasFileSystemInventoryView AliyunNasFileSystem
type AliyunNasFileSystemInventoryView struct {
	BaseInfoView
	BaseTimeView
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	StorageType string `json:"storageType,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	FileSystemId string `json:"fileSystemId,omitempty"`
}

// AddAliyunNasFileSystemEventView AddAliyunNasFileSystemEvent
type AddAliyunNasFileSystemEventView struct {
	Inventory AliyunNasFileSystemInventoryView `json:"inventory,omitempty"`
}

// CreateNasFileSystemEventView CreateNasFileSystemEvent
type CreateNasFileSystemEventView struct {
	Inventory NasFileSystemInventoryView `json:"inventory,omitempty"`
}

