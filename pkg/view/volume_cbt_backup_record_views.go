// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeCbtBackupRecordInventoryView VolumeCbtBackupRecord
type VolumeCbtBackupRecordInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"taskUuid,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"mode,omitempty"`
	rest string `json:"target,omitempty"`
	rest string `json:"scratchNodeName,omitempty"`
	rest string `json:"bitmapName,omitempty"`
	rest string `json:"lastBitmapName,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

