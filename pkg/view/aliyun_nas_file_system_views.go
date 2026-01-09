// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasFileSystemInventoryView AliyunNasFileSystem
type AliyunNasFileSystemInventoryView struct {
	DataCenterUuid *string `json:"dataCenterUuid,omitempty"`
	StorageType *string `json:"storageType,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Type *string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	FileSystemId *string `json:"fileSystemId,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// AddAliyunNasFileSystemEventView AddAliyunNasFileSystemEvent
type AddAliyunNasFileSystemEventView struct {
	Inventory AliyunNasFileSystemInventoryView `json:"inventory,omitempty"`
}

