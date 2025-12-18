// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImagePackageInventoryView ImagePackage
type ImagePackageInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	State string `json:"state,omitempty"`
	ExportUrl string `json:"exportUrl,omitempty"`
	Md5Sum string `json:"md5Sum,omitempty"`
	Format string `json:"format,omitempty"`
	Size int64 `json:"size,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

