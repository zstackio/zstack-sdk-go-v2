// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BuildAppExportHistoryInventoryView BuildAppExportHistory
type BuildAppExportHistoryInventoryView struct {
	Id int64 `json:"id,omitempty"`
	BuildAppUuid string `json:"buildAppUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Size int64 `json:"size,omitempty"`
	Md5Sum string `json:"md5Sum,omitempty"`
	Version string `json:"version,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

