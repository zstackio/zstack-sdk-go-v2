// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ZBoxBackupInventoryView ZBoxBackup
type ZBoxBackupInventoryView struct {
	ZBoxUuid string `json:"zBoxUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

