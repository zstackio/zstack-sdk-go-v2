// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BuildAppExportHistoryInventoryView BuildAppExportHistory
type BuildAppExportHistoryInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"buildAppUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"path,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"md5Sum,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

