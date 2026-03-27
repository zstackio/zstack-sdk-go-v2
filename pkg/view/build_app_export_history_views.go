// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BuildAppExportHistoryInventoryView BuildAppExportHistory
type BuildAppExportHistoryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	BuildAppUuid string `json:"buildAppUuid,omitempty"`
	Path string `json:"path,omitempty"`
	Size int64 `json:"size,omitempty"`
	Md5Sum string `json:"md5Sum,omitempty"`
	Version string `json:"version,omitempty"`
	Status string `json:"status,omitempty"`
}

// ExportBuildAppEventView ExportBuildAppEvent
type ExportBuildAppEventView struct {
	Inventory BuildAppExportHistoryInventoryView `json:"inventory,omitempty"`
}

// DeleteBuildAppExportHistoryEventView DeleteBuildAppExportHistoryEvent
type DeleteBuildAppExportHistoryEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryBuildAppExportHistoryView QueryBuildAppExportHistory
type QueryBuildAppExportHistoryView struct {
	Inventories []BuildAppExportHistoryInventoryView `json:"inventories,omitempty"`
}

