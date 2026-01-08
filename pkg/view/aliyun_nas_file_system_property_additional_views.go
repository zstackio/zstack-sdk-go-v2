// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasFileSystemPropertyView AliyunNasFileSystemProperty
type AliyunNasFileSystemPropertyView struct {
	FileSystemId string `json:"fileSystemId,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	StorageType  string `json:"storageType,omitempty"`
	Description  string `json:"description,omitempty"`
	CreateDate   string `json:"createDate,omitempty"`
}
