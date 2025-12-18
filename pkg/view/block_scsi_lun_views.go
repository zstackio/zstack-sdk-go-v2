// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BlockScsiLunInventoryView BlockScsiLun
type BlockScsiLunInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"wwn,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest int `json:"id,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"target,omitempty"`
	rest int `json:"lunMapId,omitempty"`
	rest string `json:"lunType,omitempty"`
	rest int `json:"lunInitSnapshotID,omitempty"`
	rest int64 `json:"usedSize,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

